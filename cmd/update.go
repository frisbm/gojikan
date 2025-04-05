package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
	"unicode"
)

const (
	apiDocUrl  = "https://raw.githubusercontent.com/jikan-me/jikan-rest/master/storage/api-docs/api-docs.json"
	specPath   = "api-docs.json"
	outputFile = "gojikan.go"
)

type specReplacement struct {
	old string
	new string
}

var specReplacements = []specReplacement{
	{"\"type\": \"float\"", "\"type\": \"number\""},
	{"\"maximum\": \"5\"", "\"maximum\": 5"},
	{"\"minimum\": \"1\"", "\"minimum\": 1"},
}

func main() {
	ctx := context.Background()
	fmt.Println("Starting OpenAPI spec update and code generation...")
	err := update(ctx)
	if err != nil {
		log.Fatalf("Update process failed: %v", err)
	}
	fmt.Println("Update process completed successfully.")
}
func update(ctx context.Context) (ferr error) {
	var specBuffer bytes.Buffer
	if err := pullSpec(ctx, &specBuffer); err != nil {
		return fmt.Errorf("failed to pull spec: %w", err)
	}
	modifiedSpec := replaceInSpec(specBuffer.String(), specReplacements)
	if err := writeSpec(modifiedSpec); err != nil {
		return fmt.Errorf("failed to write spec: %w", err)
	}
	if err := generate(); err != nil {
		return fmt.Errorf("code generation failed: %w", err)
	}
	if err := fmtCode(); err != nil {
		return fmt.Errorf("code formatting failed: %w", err)
	}
	if err := rewriteAst(); err != nil {
		return fmt.Errorf("AST rewrite failed: %w", err)
	}
	return nil
}
func isAnonymousStructType(typeExpr ast.Expr) bool {
	switch T := typeExpr.(type) {
	case *ast.StructType:
		return T.Fields != nil && len(T.Fields.List) > 0
	case *ast.StarExpr:
		return isAnonymousStructType(T.X)
	case *ast.ArrayType:
		return isAnonymousStructType(T.Elt)
	default:
		return false
	}
}

type ExtractionInfo struct {
	FieldPathToGeneratedType map[string]string
	GeneratedDecls           map[string]*ast.GenDecl
	GeneratedNames           map[string]struct{}
}

func NewExtractionInfo() *ExtractionInfo {
	return &ExtractionInfo{
		FieldPathToGeneratedType: make(map[string]string),
		GeneratedDecls:           make(map[string]*ast.GenDecl),
		GeneratedNames:           make(map[string]struct{}),
	}
}
func capitalize(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
func recursiveExtract(
	typeExpr ast.Expr,
	parentPath string,
	fieldName string,
	info *ExtractionInfo,
) ast.Expr {
	currentPath := parentPath
	if fieldName != "" {
		currentPath = parentPath + "." + capitalize(fieldName)
	}
	switch T := typeExpr.(type) {
	case *ast.StarExpr:
		newX := recursiveExtract(T.X, parentPath, fieldName, info)
		if newX != T.X {
			return &ast.StarExpr{X: newX}
		}
		return T
	case *ast.ArrayType:
		newElt := recursiveExtract(T.Elt, parentPath, fieldName+"Element", info)
		if newElt != T.Elt {
			return &ast.ArrayType{Elt: newElt}
		}
		return T
	case *ast.StructType:
		if !isAnonymousStructType(T) {
			return T
		}
		baseName := strings.ReplaceAll(currentPath, ".", "")
		typeName := baseName
		counter := 1
		for _, exists := info.GeneratedNames[typeName]; exists; _, exists = info.GeneratedNames[typeName] {
			typeName = fmt.Sprintf("%s%d", baseName, counter)
			counter++
		}
		info.GeneratedNames[typeName] = struct{}{}
		directAccessPath := parentPath + "." + capitalize(fieldName)
		info.FieldPathToGeneratedType[directAccessPath] = typeName
		newStructDef := &ast.StructType{
			Fields: &ast.FieldList{
				Opening: T.Fields.Opening,
				Closing: T.Fields.Closing,
			},
		}
		newTypeSpec := &ast.TypeSpec{
			Name: ast.NewIdent(typeName),
			Type: newStructDef,
		}
		newGenDecl := &ast.GenDecl{
			Tok:   token.TYPE,
			Specs: []ast.Spec{newTypeSpec},
		}
		if _, exists := info.GeneratedDecls[typeName]; !exists {
			info.GeneratedDecls[typeName] = newGenDecl
		}
		newFields := make([]*ast.Field, 0, len(T.Fields.List))
		for i, field := range T.Fields.List {
			newField := *field
			subFieldName := ""
			if len(field.Names) > 0 {
				subFieldName = field.Names[0].Name
			} else {
				subFieldName = fmt.Sprintf("Field%d", i)
			}
			newFieldType := recursiveExtract(field.Type, typeName, subFieldName, info)
			newField.Type = newFieldType
			newFields = append(newFields, &newField)
		}
		newStructDef.Fields.List = newFields
		return ast.NewIdent(typeName)
	default:
		return T
	}
}
func pass2UpdateFunctions(parsedFile *ast.File, info *ExtractionInfo) {
	ast.Inspect(parsedFile, func(node ast.Node) bool {
		funcDecl, isFunc := node.(*ast.FuncDecl)
		if !isFunc || funcDecl.Body == nil {
			return true
		}
		localVarAnonStructDefNodes := make(map[string]ast.Node)
		varTypes := make(map[string]string)
		ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
			switch stmt := n.(type) {
			case *ast.AssignStmt:
				switch stmt.Tok {
				case token.DEFINE:
					if len(stmt.Lhs) == 1 && len(stmt.Rhs) == 1 {
						if ident, ok := stmt.Lhs[0].(*ast.Ident); ok {
							varName := ident.Name
							if ue, ok := stmt.Rhs[0].(*ast.UnaryExpr); ok && ue.Op == token.AND {
								if cl, ok := ue.X.(*ast.CompositeLit); ok {
									if typeIdent, ok := cl.Type.(*ast.Ident); ok {
										varTypes[varName] = typeIdent.Name
									}
									if st, ok := cl.Type.(*ast.StructType); ok && isAnonymousStructType(st) {
										localVarAnonStructDefNodes[varName] = stmt
									}
								}
							}
							if cl, ok := stmt.Rhs[0].(*ast.CompositeLit); ok {
								if typeIdent, ok := cl.Type.(*ast.Ident); ok {
									varTypes[varName] = typeIdent.Name
								}
							}
						}
					}
				case token.ASSIGN:
					if len(stmt.Lhs) == 1 && len(stmt.Rhs) == 1 {
						selExpr, isSel := stmt.Lhs[0].(*ast.SelectorExpr)
						unaryExpr, isUnary := stmt.Rhs[0].(*ast.UnaryExpr)
						if isSel && isUnary && unaryExpr.Op == token.AND {
							if rhsIdent, ok := unaryExpr.X.(*ast.Ident); ok {
								localVarName := rhsIdent.Name
								if declNode, exists := localVarAnonStructDefNodes[localVarName]; exists {
									targetVarName := ""
									if targetIdent, ok := selExpr.X.(*ast.Ident); ok {
										targetVarName = targetIdent.Name
									}
									if targetTypeName, typeFound := varTypes[targetVarName]; typeFound {
										fieldName := capitalize(selExpr.Sel.Name)
										lookupKey := targetTypeName + "." + fieldName
										if generatedTypeName, genTypeFound := info.FieldPathToGeneratedType[lookupKey]; genTypeFound {
											replaced := false
											switch decl := declNode.(type) {
											case *ast.AssignStmt:
												if ue, ok := decl.Rhs[0].(*ast.UnaryExpr); ok {
													if cl, ok := ue.X.(*ast.CompositeLit); ok {
														cl.Type = ast.NewIdent(generatedTypeName)
														replaced = true
													}
												}
											case *ast.ValueSpec:
												if len(decl.Values) == 1 {
													if ue, ok := decl.Values[0].(*ast.UnaryExpr); ok && ue.Op == token.AND {
														if cl, ok := ue.X.(*ast.CompositeLit); ok {
															cl.Type = ast.NewIdent(generatedTypeName)
															replaced = true
														}
													}
												}
												if !replaced {
													if st, ok := decl.Type.(*ast.StructType); ok && isAnonymousStructType(st) {
														decl.Type = ast.NewIdent(generatedTypeName)
														replaced = true
													}
												}
											}
											if replaced {
												delete(localVarAnonStructDefNodes, localVarName)
											}
										}
									}
								}
							}
						}
						if isSel && isUnary && unaryExpr.Op == token.AND {
							if compLit, ok := unaryExpr.X.(*ast.CompositeLit); ok {
								if _, isAnonStruct := compLit.Type.(*ast.StructType); isAnonStruct && isAnonymousStructType(compLit.Type) {
									targetVarName := ""
									if targetIdent, ok := selExpr.X.(*ast.Ident); ok {
										targetVarName = targetIdent.Name
									}
									if targetTypeName, typeFound := varTypes[targetVarName]; typeFound {
										fieldName := capitalize(selExpr.Sel.Name)
										lookupKey := targetTypeName + "." + fieldName
										if generatedTypeName, genTypeFound := info.FieldPathToGeneratedType[lookupKey]; genTypeFound {
											compLit.Type = ast.NewIdent(generatedTypeName)
										}
									}
								}
							}
						}
					}
				}
			case *ast.GenDecl:
				if stmt.Tok == token.VAR {
					for _, spec := range stmt.Specs {
						if valSpec, ok := spec.(*ast.ValueSpec); ok {
							if len(valSpec.Names) == 1 {
								varName := valSpec.Names[0].Name
								if valSpec.Type != nil {
									if typeIdent, ok := valSpec.Type.(*ast.Ident); ok {
										varTypes[varName] = typeIdent.Name
									}
									if starExpr, ok := valSpec.Type.(*ast.StarExpr); ok {
										if typeIdent, ok := starExpr.X.(*ast.Ident); ok {
											varTypes[varName] = typeIdent.Name
										}
									}
									if st, ok := valSpec.Type.(*ast.StructType); ok && isAnonymousStructType(st) {
										localVarAnonStructDefNodes[varName] = valSpec
									}
								}
								if len(valSpec.Values) == 1 {
									if ue, ok := valSpec.Values[0].(*ast.UnaryExpr); ok && ue.Op == token.AND {
										if cl, ok := ue.X.(*ast.CompositeLit); ok {
											if typeIdent, ok := cl.Type.(*ast.Ident); ok {
												varTypes[varName] = typeIdent.Name
											}
											if st, ok := cl.Type.(*ast.StructType); ok && isAnonymousStructType(st) {
												localVarAnonStructDefNodes[varName] = valSpec
											}
										}
									}
									if cl, ok := valSpec.Values[0].(*ast.CompositeLit); ok {
										if typeIdent, ok := cl.Type.(*ast.Ident); ok {
											varTypes[varName] = typeIdent.Name
										}
									}
								}
							}
						}
					}
				}
			}
			return true
		})
		return false
	})
}
func rewriteAst() (err error) {
	fset := token.NewFileSet()
	parsedFile, err := parser.ParseFile(fset, outputFile, nil, 0)
	if err != nil {
		return fmt.Errorf("failed to parse file '%s': %w", outputFile, err)
	}
	extractionInfo := NewExtractionInfo()
	ast.Inspect(parsedFile, func(n ast.Node) bool {
		if ts, ok := n.(*ast.TypeSpec); ok {
			extractionInfo.GeneratedNames[ts.Name.Name] = struct{}{}
		}
		return true
	})
	ast.Inspect(parsedFile, func(node ast.Node) bool {
		typeSpec, isTypeSpec := node.(*ast.TypeSpec)
		if !isTypeSpec {
			return true
		}
		structType, isStruct := typeSpec.Type.(*ast.StructType)
		if !isStruct || structType.Fields == nil {
			return true
		}
		parentTypeName := typeSpec.Name.Name
		if structType.Fields != nil {
			for _, field := range structType.Fields.List {
				if len(field.Names) > 0 {
					fieldName := field.Names[0].Name
					newFieldType := recursiveExtract(field.Type, parentTypeName, fieldName, extractionInfo)
					if newFieldType != field.Type {
						field.Type = newFieldType
					}
				}
			}
		}
		return false
	})
	pass2UpdateFunctions(parsedFile, extractionInfo)
	newDeclsSlice := make([]ast.Decl, 0, len(extractionInfo.GeneratedDecls))
	for _, decl := range extractionInfo.GeneratedDecls {
		newDeclsSlice = append(newDeclsSlice, decl)
	}
	insertIndex := -1
	for i, decl := range parsedFile.Decls {
		if gd, ok := decl.(*ast.GenDecl); ok && gd.Tok == token.TYPE {
			isGenerated := false
			if ts, ok := gd.Specs[0].(*ast.TypeSpec); ok {
				if _, exists := extractionInfo.GeneratedDecls[ts.Name.Name]; exists {
					isGenerated = true
				}
			}
			if !isGenerated {
				insertIndex = i
				break
			}
		}
	}
	if insertIndex != -1 {
		parsedFile.Decls = append(parsedFile.Decls[:insertIndex], append(newDeclsSlice, parsedFile.Decls[insertIndex:]...)...)
	} else {
		parsedFile.Decls = append(parsedFile.Decls, newDeclsSlice...)
	}
	var buf bytes.Buffer
	err = format.Node(&buf, fset, parsedFile)
	if err != nil {
		return fmt.Errorf("failed to format rewritten code: %w", err)
	}
	err = os.WriteFile(outputFile, buf.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("failed to write modified file '%s': %w", outputFile, err)
	}
	return nil
}
func fmtCode() (ferr error) {
	cmd := exec.Command("make", "fmt")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		errMsg := fmt.Sprintf("'make fmt' failed with error: %v", err)
		if stderrStr := stderr.String(); stderrStr != "" {
			errMsg += fmt.Sprintf("\nstderr:\n%s", stderrStr)
		}
		return errors.New(errMsg)
	}
	return nil
}
func generate() (ferr error) {
	cmd := exec.Command("go", "generate", "./...")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		errMsg := fmt.Sprintf("'go generate ./...' failed with error: %v", err)
		if stderrStr := stderr.String(); stderrStr != "" {
			errMsg += fmt.Sprintf("\nstderr:\n%s", stderrStr)
		}
		return errors.New(errMsg)
	}
	return nil
}
func writeSpec(spec string) (ferr error) {
	if err := os.WriteFile(specPath, []byte(spec), 0644); err != nil {
		return fmt.Errorf("failed to write spec file '%s': %w", specPath, err)
	}
	return nil
}
func replaceInSpec(spec string, replacements []specReplacement) string {
	result := spec
	for _, replacement := range replacements {
		result = strings.ReplaceAll(result, replacement.old, replacement.new)
	}
	return result
}
func pullSpec(ctx context.Context, buf io.Writer) (ferr error) {
	client := httpClient()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiDocUrl, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return fmt.Errorf("HTTP request timed out: %w", err)
		}
		if errors.Is(err, context.Canceled) {
			return fmt.Errorf("HTTP request cancelled: %w", err)
		}
		return fmt.Errorf("failed to execute HTTP request: %w", err)
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			ferr = errors.Join(ferr, fmt.Errorf("failed to close response body: %w", cerr))
		}
	}()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected HTTP status: %s", resp.Status)
	}
	if _, err = io.Copy(buf, resp.Body); err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}
	return nil
}
func httpClient() *http.Client {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   15 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout:   15 * time.Second,
		ResponseHeaderTimeout: 15 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConns:          10,
		IdleConnTimeout:       60 * time.Second,
	}
	client := &http.Client{
		Transport: transport,
		Timeout:   60 * time.Second,
	}
	return client
}

# GoJikan

[![Go Report Card](https://goreportcard.com/badge/github.com/frisbm/gojikan)](https://goreportcard.com/report/github.com/frisbm/gojikan)
[![Go Reference](https://pkg.go.dev/badge/github.com/frisbm/gojikan.svg)](https://pkg.go.dev/github.com/frisbm/gojikan)

## ⚠️ Project Status: Under Development ⚠️

GoJikan is a Go client library for the [Jikan API v4](https://docs.api.jikan.moe/), which is an unofficial REST API
for [MyAnimeList](https://myanimelist.net). It provides a simple interface for Go developers to fetch data about anime,
manga, characters, and more from MyAnimeList via the Jikan API.

## Features

* Easy-to-use Go interface for Jikan API v4 endpoints.
* Provides Go structs for Jikan API responses.
* Aims to cover the full Jikan API v4 surface.
* Client Configuration for custom settings including custom base URL, HTTP client, and cache.

### Installation

To add GoJikan to your Go project:

```sh
  go get -u github.com/frisbm/gojikan
```

### FAQ:

* **What is Jikan?**
    * Jikan is a RESTful API for accessing data from MyAnimeList. It provides a simple and easy-to-use interface for
      developers to interact with the MyAnimeList API and retrieve information about anime, manga, characters, and more.
* **What is GoJikan?**
    * GoJikan is a Go client for the Jikan API. It provides a simple and easy-to-use interface for developers to
      interact with the Jikan API and retrieve information about anime, manga, characters, and more.
* **Why did I handroll my own client?**
    * The OpenAPI generators I tried weren't playing nicely with the Jikan OpenAPI spec. I tried a few different ones,
      but they all had issues with the spec. So I decided to handroll my own client instead.


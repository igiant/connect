# Kerio Connect API
[![Go Reference](https://pkg.go.dev/badge/github.com/igiant/connect.svg)](https://pkg.go.dev/github.com/igiant/connect)
## Overview
Client for [Kerio API Connect (JSON-RPC 2.0)](https://manuals.gfi.com/en/kerio/api/connect/admin/reference/index.html)

Implemented several Administration API for Kerio Connect methods

Created the basis for easily adding your own methods

## Installation
```go
go get github.com/igiant/connect
```

## Example
```go
package main

import (
	"fmt"
	"log"

	"github.com/igiant/connect"
)

func main() {
	config := connect.NewConfig(connect.Admin, "server_addr")
	conn, err := config.NewConnection()
	if err != nil {
		log.Fatal(err)
	}
	app := &connect.Application{
		Name:    "MyApp",
		Vendor:  "Me",
		Version: "v0.0.1",
	}
	err = conn.Login("user_name", "user_password", app)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = conn.Logout()
		if err != nil {
			log.Println(err)
		}
	}()
	info, err := conn.ServerGetProductInfo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(
		"ProductName: %s\nVersion: %s\nOsName: %s\n",
		info.ProductName,
		info.Version,
		info.OsName,
	)
}
```
## Documentation
* [GoDoc](http://godoc.org/github.com/igiant/connect)

## RoadMap
* Add remaining methods for Kerio Connect

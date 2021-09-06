# godb-go-driver
A [GoDB](https://github.com/MattLaidlaw/godb) client for the Go language. This package allows a Go program to interact with a GoDB server via TCP connection.

## Requirements
* Go 1.17

## Install
```
go get github.com/MattLaidlaw/GoDB-Go-Driver@v1.0.0
```

## Usage
The below example shows the creation of a GoDB client and the possible methods it can perform.
```go
package main

import (
	"fmt"
	"github.com/MattLaidlaw/godb-go-driver/pkg/driver"
	"log"
)

func main() {

	// create a godb client
	client, err := driver.NewClient("localhost:6342")
	if err != nil {
		log.Fatalln(err)
	}

	// set a key-value pair in the database
	insertedCount, err := client.Set("key", "godb")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(insertedCount)  // expect insertedCount = 1

	// get a key-value pair by key
	value, err := client.Get("key")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(value)  // expect value = "godb"

	// delete a key-value pair by key
	deletedCount, err := client.Del("key")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(deletedCount)  // expect deletedCount = 1

}
```

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
  "github.com/MattLaidlaw/GoDB-Go-Driver/pkg/driver"
  "log"
)

func main() {

  // create a godb client
  client, err := driver.NewClient("localhost:6342")
  if err != nil {
    log.Fatalln(err)
  }
  
  // set a key-value pair in the database
  insertedCount, err := client.Set("key", "godb")  // expect insertedCount = 1
  if err != nil {
    log.Println(err)
  }
  
  // get a key-value pair by key
  value, err := client.Get("key")  // expect value = "godb"
  if err != nil {
    log.Println(err)
  }
  
  // delete a key-value pair by key
  deletedCount, err := client.Del("key")  // expect deletedCount = 1
  if err != nil {
    log.Println(err)
  }
  
}
```

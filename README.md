# GoodiesDB Client Lib
[![Go](https://github.com/andrelcunha/GoodiesDB-client-lib/actions/workflows/go.yml/badge.svg)](https://github.com/andrelcunha/GoodiesDB-client-lib/actions/workflows/go.yml)

### Overview
GoodiesDB Client Lib is a Go client library for interacting with the [GoodiesDB](https://github.com/andrelcunha/GoodiesDB) server, a Redis-like key-value store. This library provides functionality to connect to the GoodiesDB server, authenticate, and execute common commands such as `SET` and `GET`.
### Features
- Connect to GoodiesDB server
- Authenticate with the server
- Execute `SET` and `GET` commands
- Support for URL-based and parameter-based client initialization
### Usage
### Client Initialization
You can initialize the client using either a URL with embedded credentials or separate address and password fields.
#### Using URL
```GO
package main

import (
    "fmt"
    "log"
    "github.com/andrelcunha/GoodiesDB-client-lib/goodiesdb"
)

func main() {
    client, err := goodiesdb.NewClientByUrl("http://user:password@localhost:6379")
    if err != nil {
        log.Fatalf("Failed to create client: %v", err)
    }

    err = client.Set("mykey", "myvalue")
    if err != nil {
        log.Fatalf("Failed to set value: %v", err)
    }

    value, err := client.Get("mykey")
    if err != nil {
        log.Fatalf("Failed to get value: %v", err)
    }

    fmt.Println("Value:", value)
}

```
#### Using Addres and Password
```Go
package main

import (
    "fmt"
    "log"
    "github.com/andrelcunha/GoodiesDB-client-lib/goodiesdb"
)

func main() {
    client, err := goodiesdb.NewClient("localhost:6379", "mypassword")
    if err != nil {
        log.Fatalf("Failed to create client: %v", err)
    }

    err = client.Set("mykey", "myvalue")
    if err != nil {
        log.Fatalf("Failed to set value: %v", err)
    }

    value, err := client.Get("mykey")
    if err != nil {
        log.Fatalf("Failed to get value: %v", err)
    }

    fmt.Println("Value:", value)
}
```
### Functons
#### Auth 
```Go
func (c *Client) Auth(password string) error
```
#### Set
```Go
func (c *Client) Set(key, value string) error
```
#### Get
```Go
func (c *Client) Get(key string) (string, error)
```
### Contributing
Contributions are welcome! Please submit a pull request or create an issue to discuss your ideas.
### License
This project is licensed under the MIT License. See the ￼LICENSE file for details.

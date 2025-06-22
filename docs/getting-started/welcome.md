# About
Goblox is a Go library that makes accessing Roblox's OpenCloud & Legacy APIs with your Go projects easy.

It supports ALL OpenCloud V2 APIs with V1 OpenCloud APIs & Legacy APIs coming sometime in the future.

## Installation
::: info
Go 1.18 or newer is required.
:::

```bash
go get -u github.com/typical-developers/goblox
```

## Basic Usage
```go
package main

import (
    "context"
    "fmt"

    "github.com/typical-developers/goblox/opencloud"
)

func main() {
    client := opencloud.NewClientWithAPIKey("YOUR_API_KEY")
}
```
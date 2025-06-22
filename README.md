# goblox
Goblox is a Go library for accessing Roblox's Open Cloud & Legacy APIs.

> [!CAUTION]
> This library is still in development and is not considered stable.
> If you want to contribute, feel free to open a PR.

## ☁️ Opencloud Basic Usage
More in-depth documentation will be added sometime in the future.

### With API Key
```go
package main

import (
    "context"
    "fmt"

    "github.com/typical-developers/goblox/opencloud"
)

func main() {
    ctx := context.Background()
    client := opencloud.NewClientWithAPIKey("YOUR_API_KEY")

    user, resp, err := client.UserandGroups.GetUser(ctx, "USER_ID")
    if err != nil {
        panic(err)
    }

    fmt.Println(user, resp.StatusCode)
}
```

### With OAuth Token
```go
package main

import (
    "context"
    "fmt"

    "github.com/typical-developers/goblox/opencloud"
)

func main() {
    ctx := context.Background()
    client := opencloud.NewClientWithOAuth("YOUR_OAUTH_TOKEN")

    user, resp, err := client.UserandGroups.GetUser(ctx, "USER_ID")
    if err != nil {
        panic(err)
    }

    fmt.Println(user, resp.StatusCode)
}
```

## 📥 Legacy API Basic Usage
Legacy API support will be added in the future.
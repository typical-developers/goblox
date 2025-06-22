# Goblox
Goblox is a Go library for accessing the Roblox Open Cloud API & Roblox Legacy API.

> [!CAUTION]
> This library is still in development and is not considered stable.
> If you want to contribute, feel free to open a PR.

## Basic Usage
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

    user, resp, err := client.UserandGroups.GetUser(ctx, "UNIVERSE_ID", "USER_ID")
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

    user, resp, err := client.UserandGroups.GetUser(ctx, "UNIVERSE_ID", "USER_ID")
    if err != nil {
        panic(err)
    }

    fmt.Println(user, resp.StatusCode)
}
```
# goblox <img src="docs/public/goblox_icon.png" align="right" width="150" alt="Goblox logo" />

Easily integrate Roblox's OpenCloud & legacy APIs into your Go projects.

> [!CAUTION]
> While this library is used in some production deployments, it is not considered fully stable. If you want to contribute, feel free to open a PR!

[![Go Reference](https://pkg.go.dev/badge/github.com/typical-developers/goblox.svg)](https://pkg.go.dev/github.com/typical-developers/goblox)
[![Go Version](https://img.shields.io/github/go-mod/go-version/typical-developers/goblox)](go.mod)
[![License: GPL v3](https://img.shields.io/badge/license-GPLv3-blue.svg)](LICENSE)

**[Documentation](https://goblox.dev)** · **[OpenCloud Guide](https://goblox.dev/guides/opencloud/authentication)** · **[Legacy Guide](https://goblox.dev/guides/legacy/authentication)**

### Features
- Support for all OpenCloud APIs.
- Easily authenticate with cookies, api keys, and oauth2 tokens.
    - CSRF tokens automatically get handled when using cookies.
    - Authentication is checked per-endpoint to prevent the wrong credential types from being used.
    - API keys and Oauth2 tokens optionally support being introspected to prevent invalid requests from being sent from missing scopes.
- Automatic rate limit retrying, enabled by default.

## Installation
```sh
go get -u github.com/typical-developers/goblox
```
### Quick Start (OpenCloud)
```go
package main

import (
    "context"
    "fmt"

    "github.com/typical-developers/goblox/opencloud"
)

func main() {
    client := opencloud.NewClient().WithAPIKey("YOUR_API_KEY")

    user, _, err := client.Users.GetUser(context.TODO(), "1")
    if err != nil {
        panic(err)
    }

    fmt.Printf("User: %s\n", user.Name)
}
```

## Current Coverage
### `opencloud`
| Service | Status | | Service | Status |
|---|:-:|---|---|:-:|
| Ads | ✅ | | Luau Execution | ✅ |
| Analytics | ✅ | | Matchmaking | ✅ |
| Assets | ✅ | | Monetization | ✅ |
| Badges | ✅ | | Places | ✅ |
| Config | ✅ | | Secrets | ✅ |
| Creator Store | ✅ | | Server Management | ✅ |
| Data & Memory Stores | ✅ | | Service Efficiency | ✅ |
| Events | ✅ | | Thumbnails | ✅ |
| Groups | ✅ | | Universes | ✅ |
| Localization | ✅ | | Users | ✅ |
### `legacy`
Contributions are welcome!
| Service | Status | | Service | Status |
|---|:-:|---|---|:-:|
| Account Information | 🚧 | | Groups Moderation | 🚧 |
| Account Settings | 🚧 | | Inventory | 🚧 |
| Ad Configuration | 🚧 | | Item Configuration | 🚧 |
| Ads | 🚧 | | Locale | 🚧 |
| Asset Delivery | 🚧 | | Localization Tables | 🚧 |
| Auth | 🚧 | | Notifications | 🚧 |
| Avatar | 🚧 | | Premium Features | 🚧 |
| Badges | 🚧 | | Presence | 🚧 |
| Billing | 🚧 | | Private Messages | 🚧 |
| Catalog | 🚧 | | Publish | 🚧 |
| Chat | 🚧 | | Search | 🚧 |
| Contacts | 🚧 | | Text Filter | 🚧 |
| Develop | 🚧 | | Thumbnails | ✅ |
| Economy | 🚧 | | Trades | 🚧 |
| Economy Creator Stats | 🚧 | | Translation Roles | 🚧 |
| Engagement Payouts | 🚧 | | Translations | 🚧 |
| Followings | 🚧 | | Two-Step Verification | 🚧 |
| Friends | 🚧 | | User Moderation | 🚧 |
| Game Internationalization | 🚧 | | Users | 🚧 |
| Game Join | 🚧 | | Voice | 🚧 |
| Game Persistence | 🚧 | | | |
| Games | 🚧 | | | |
| Groups | 🚧 | | | |
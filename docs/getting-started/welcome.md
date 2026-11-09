# About <img src="/goblox_icon.png" align="right" width="150" />

Easily integrate Roblox's OpenCloud & legacy APIs into your Go projects.

<div class="badges">

[![Go Reference](https://pkg.go.dev/badge/github.com/luckfire/goblox.svg)](https://pkg.go.dev/github.com/luckfire/goblox)
[![Go Version](https://img.shields.io/github/go-mod/go-version/luckfire/goblox)](https://github.com/luckfire/goblox/blob/main/go.mod)
[![License: GPL v3](https://img.shields.io/badge/license-GPLv3-blue.svg)](https://github.com/luckfire/goblox/blob/main/LICENSE)

</div>

### Features
- Support for all OpenCloud APIs.
- Easily authenticate with cookies, api keys, and oauth2 tokens.
    - CSRF tokens automatically get handled when using cookies.
    - Authentication is checked per-endpoint to prevent the wrong credential types from being used.
    - API keys and Oauth2 tokens optionally support being introspected to prevent invalid requests from being sent from missing scopes.
- Automatic rate limit retrying, enabled by default.

## Installation
::: info
Go 1.23 or newer is required.
:::

```bash
go get -u github.com/luckfire/goblox
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
| Service | Status | | Service | Status |
|---|:-:|---|---|:-:|
| Account Information | 🚧 | | Groups Moderation | 🚧 |
| Account Settings | 🚧 | | Inventory | 🚧 |
| Ad Configuration | 🚧 | | Item Configuration | 🚧 |
| Ads | 🚧 | | Locale | 🚧 |
| Asset Delivery | 🚧 | | Localization Tables | 🚧 |
| Auth | 🚧 | | Notifications | 🚧 |
| Avatar | 🚧 | | Premium Features | 🚧 |
| Badges | ✅ | | Presence | ✅ |
| Billing | 🚧 | | Private Messages | 🚧 |
| Catalog | 🚧 | | Publish | 🚧 |
| Chat | 🚧 | | Search | 🚧 |
| Contacts | 🚧 | | Text Filter | 🚧 |
| Develop | 🚧 | | Thumbnails | ✅ |
| Economy | 🚧 | | Trades | 🚧 |
| Economy Creator Stats | 🚧 | | Translation Roles | 🚧 |
| Engagement Payouts | 🚧 | | Translations | 🚧 |
| Followings | 🚧 | | Two-Step Verification | 🚧 |
| Friends | ✅ | | User Moderation | 🚧 |
| Game Internationalization | 🚧 | | Users | ✅ |
| Game Join | 🚧 | | Voice | 🚧 |
| Game Persistence | 🚧 | | | |
| Games | ✅ | | | |
| Groups | 🚧 | | | |
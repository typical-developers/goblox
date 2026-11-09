# Authentication
::: tip
These are legacy APIs, which means API keys and OAuth2 tokens will not work. Some legacy APIs may be mirrored onto Roblox's OpenCloud endpoints. See the [OpenCloud authentication guide](/guides/opencloud/authentication).
:::
The legacy APIs on `*.roblox.com` take a `.ROBLOSECURITY` cookie. Many of their endpoints are public and need no credential at all.

## Cookies
::: danger
A `.ROBLOSECURITY` cookie will give full access to the account. It is advised to use a separate account for automation. Roblox does not recommend cookie-based endpoints for production and states that they can change without notice. You should always prefer using OpenCloud endpoints with an API key or OAuth2 token wherever possible.
:::
For endpoints that are locked behind a cookie, you can pass an account's `.ROBLOSECURITY` cookie. You do not need to include the cookie's key, `.ROBLOSECURITY=`. 

### Using Cookies (client-wide)
```go
package main

import (
    "github.com/luckfire/goblox/legacy"
)

func main() {
    client := legacy.NewClient().WithCookie("YOUR_ROBLOSECURITY_COOKIE")
}
```
#### Using Cookies (per-call)
```go
import (
    "context"

    "github.com/luckfire/goblox"
    "github.com/luckfire/goblox/auth"
)

var client = legacy.NewClient()

func main() {
    cookie := auth.NewROBLOSecurityCookie("YOUR_ROBLOSECURITY_COOKIE")
    client.Thumbnails.Batch(context.TODO(), []legacy.ThumbnailBatch{}, goblox.WithCredential(cookie))
}

```

## CSRF tokens
Roblox requires an `X-CSRF-Token` header on requests that change data (`POST`, `PATCH`, `PUT` and `DELETE`) made with a cookie. Goblox will handle getting this token for you.

## Handling Authentication Errors
These errors are returned before a request is sent to Roblox's servers.

If Roblox rejects the cookie itself (for example, because it expired or the account signed out everywhere), the request comes back with a 401.
```go
import (
    "errors"

    "github.com/luckfire/goblox"
)

var unsupported *goblox.UnsupportedAuthError
if errors.As(err, &unsupported) {
    // This endpoint needs a cookie. unsupported.Accepts lists what it takes.
}
```

# Authentication
::: tip
Some endpoints may only accept an API key or an OAuth2 token. Each method's doc comment lists the credentials and the scopes it needs.
:::
The OpenCloud APIs take two different types of credentials:

- **[API keys](https://create.roblox.com/docs/cloud/auth/api-keys)**, which you can create either a user or group api key.
- **[OAuth2 tokens](https://create.roblox.com/docs/cloud/auth/oauth2-overview)**, which users grant when they authorize with your application.

Every endpoint declares which authorization type is accepted alongside the scopes required for that endpoint. If you call one with a credential it does not take, Goblox will return an error before requesting an endpoint.

## Using API Keys
You can create an API key at: https://create.roblox.com/dashboard/credentials?activeTab=ApiKeysTab
```go
package main

import (
    "context"
    "fmt"

    "github.com/luckfire/goblox/opencloud"
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

## OAuth2 Tokens
You create an OAuth2 application at: https://create.roblox.com/dashboard/credentials?activeTab=OAuthTab
### Using OAuth2 Tokens (client-wide) 
```go
package main

import (
    "context"
    "fmt"

    "github.com/luckfire/goblox"
    "github.com/luckfire/goblox/opencloud"
)

func main() {
    authedUser := opencloud.NewClient().WithOAuthToken("OAUTH2_TOKEN")

    user, _, err := authedUser.Users.GetUser(context.TODO(), "1")
    if err != nil {
        panic(err)
    }

    fmt.Printf("User: %s\n", user.Name)
}
```
### Using OAuth2 Tokens (per-call) 
```go
package main

import (
    "context"
    "fmt"

    "github.com/luckfire/goblox"
    "github.com/luckfire/goblox/opencloud"
)

var client = opencloud.NewClient()

func main() {
    user, _, err := client.Users.GetUser(context.TODO(),
        "1", goblox.WithOAuthToken("OAUTH2_TOKEN")
    )
    if err != nil {
        panic(err)
    }

    fmt.Printf("User: %s\n", user.Name)
}
```
Per-call options like `goblox.WithOAuthToken` override the client's credential for that call only. This makes it safe to use from many goroutines.

## Introspecting Scopes
By default, Goblox will not know which scopes your key or token has. To have Goblox check first and return a clear error instead of a 403, you can introspect your credential.

The first call with an introspected credential makes one extra request to Roblox. Later calls use the cached result until it expires.
### API Keys
API credentials can be set to be checked at an interval. By default if the duration is set to 0, it will check every 10 minutes. It will also check after every 401/403 until the key scopes are updated.
```go
import (
    "time"

    "github.com/luckfire/goblox"
    "github.com/luckfire/goblox/opencloud"
    "github.com/luckfire/goblox/auth"
)

func main() {
    key := auth.IntrospectedAPIKey("YOUR_API_KEY", 10*time.Minute)
    client := opencloud.NewClient().WithCredential(key)
}
```
### OAuth2 Token
An OAuth2 credential will need your app's client ID and secret.
```go
import (
    "time"

    "github.com/luckfire/goblox"
    "github.com/luckfire/goblox/opencloud"
    "github.com/luckfire/goblox/auth"
)

func main() {
    token := auth.IntrospectedOAuthToken("USER_OAUTH2_TOKEN", "YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET") 
    client := opencloud.NewClient().WithCredential(token)
}
```

## Handling Authentication Errors
These errors are returned before a request is sent to Roblox's servers.
```go
import (
    "errors"

    "github.com/luckfire/goblox"
    "github.com/luckfire/goblox/auth"
)

var unsupported *goblox.UnsupportedAuthError
var missing *goblox.MissingScopesError

switch {
case errors.As(err, &unsupported):
    // The endpoint doesn't take this kind of credential (for example, OAuth on data stores).
case errors.As(err, &missing):
    // The credential is the right kind, but lacks a scope.
    // missing.Missing lists required scopes; missing.AnyOf lists alternatives.
case errors.Is(err, auth.ErrGrantInactive):
    // The key is disabled or expired, or the token is no longer active.
}
```
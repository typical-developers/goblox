# goblox
A Golang wrapper for the Roblox API. Supports Opencloud and Legacy APIs.

## Examples 📝
### Opencloud
```go
package main

import (
	"github.com/typical-developers/goblox/opencloud"
)

func main() {
    // Set your opencloud API token
    opencloud.SetAPIToken("YOUR_API_TOKEN")

    // In this example, we will be fetching information for a Universe.
    // Your scopes must have universe.place:write for this to work.
    universe, _ := opencloud.GetUniverse("UNIVERSE_ID")
    println(universe.DisplayName) // Name of the Universe.
}
```
### Legacy API
Legacy API is a work in progress.

## Other Features ✨
Coming soon!
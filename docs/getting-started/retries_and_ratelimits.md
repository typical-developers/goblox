# Retries & Ratelimits
::: info
The longest that Goblox will wait is up to 30 seconds before a retry. Even if Roblox gives its own time, it will still return an error instead of waiting.
:::
Enabled by default, Goblox will handle `429 Too Many Request`, or ratelimits. If possible, Goblox will wait as long as Roblox requests and send the request again.

Goblox will use a few factors to decide on how it should retry a request.
- If Roblox provides a `Retry-After` header, it will use the duration from that before retrying.
- If Roblox provides a `x-ratelimit-reset` and `x-ratelimit-remaining` header, it will calculate a ratelimit based on the reset if the remaining is 0.
- An expoential backoff (500 miliseconds, 1 second, 2 second ... and so on).

## Checking Ratelimits
Every response carries Roblox's ratelimit headers, assuming that Roblox sent any. You are able to check this ratelimit by checking the response. `Remaining` will return as `-1` when there is no ratelimit headers. Legacy APIs do not return any ratelimit headers.

API key limits are shared by every key with the same owner, user, or group. `Remaining` counts all of that owner's traffic, not only requests from the client. OAuth limits are per access token, so every user will have their own. You can read more about ratelimits in [Roblox's Documentation](https://create.roblox.com/docs/cloud/reference/rate-limits).
```go
_, resp, err := client.Users.GetUser(ctx, "USER_ID")
if err == nil {
    ratelimit := resp.RateLimit()
    fmt.Printf("%d requests left, resets in %s\n", rl.Remaining, rl.Reset)
}
```

## Setting Retries & Policies
Both `opencloud` and `legacy` clients let you customize the retry policy. You can either set a custom client-wide retry policy, toggle retries on for the client without messing with the policy, or toggle retries per endpoint.
### Policy
```go
opencloud.NewClient(opencloud.WithRetryPolicy(goblox.RetryPolicy{
    MaxRetries: 5,
}))
legacy.NewClient(legacy.WithRetryPolicy(goblox.RetryPolicy{
    MaxRetries: 5,
}))
```
### Toggling
```go
opencloud.NewClient(opencloud.WithRetries(false))
legacy.NewClient(legacy.WithRetries(false))
```
```go
_, resp, err := client.Users.GetUser(ctx, "USER_ID", goblox.WithRetries(false))
```
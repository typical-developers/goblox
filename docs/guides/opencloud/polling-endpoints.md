# Polling Endpoints
In some circumstances, such as the Luau Execution APIs, you are expected to poll (or request these endpoints over an interval) until you get the desired result. Below is the example taken from the [Luau Execution](/guides/opencloud/luau-execution.html) page.

```go
package main

import (
    "context"
    "fmt"

    "github.com/typical-developers/goblox/opencloud"
)

func main() {
    ctx := context.Background()
    client := opencloud.NewClient().WithAPIKey("YOUR_API_KEY")

    // First, we create the task with the Luau execution API.
    task, _, err := client.LuauExecution.CreateLuauExecutionSessionTask(ctx, "UNIVERSE_ID", "PLACE_ID", nil, opencloud.LuauExecutionTaskCreate{
        Script: opencloud.Pointer("return 1 + 2"),
    })
    if err != nil {
        panic(err)
    }

    // Then, we get the TaskInfo so we can get the task we just created.
    universeID, placeID, versionId, sessionId, taskId := task.TaskInfo()

    // Finally, we poll the task until it's complete and set the reuslt / error in our variables.
    var results []any
    var err error

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			task, resp, err := client.LuauExecution.GetLuauExecutionSessionTask(ctx, universeID, placeID, versionId, sessionId, taskId)

			if err != nil {
				return nil, err
			}

			if resp.StatusCode == http.StatusTooManyRequests {
				log.Warn("LuauExecution is being ratelimited.")
				continue
			}
			if task.State == opencloud.LuauExecutionStateQueued || task.State == opencloud.LuauExecutionStateProcessing {
				continue
			}

			if task.Output != nil {
                results = task.Output
                break
			}

			if task.Error != nil {
                err = errors.New(task.Error.Message)
                break
			}
		case <-ctx.Done():
            err = ctx.Err()
            break
		}
	}

    fmt.Println(fmt.Sprintf("Results: %+v", results))
    fmt.Println(fmt.Sprintf("Error: %+v", err))
}
```
# Luau Execution
The OpenCloud APIs allow you to execute Luau code in one of your experiences by spinning up a experience instance and executing the code.

### Executing Luau
The code below will create a new task operation for executing the code in the experience.
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

    task, _, err := client.LuauExecution.CreateLuauExecutionSessionTask(ctx, "UNIVERSE_ID", "PLACE_ID", nil, opencloud.LuauExecutionTaskCreate{
        Script: opencloud.Pointer("return 1 + 2"),
    })
    if err != nil {
        panic(err)
    }
}
```

However, this in itself is not very useful, since it runs asynchronously. We need a way to check when the task is finished and get the results or error from it. This is where polling the method comes in.

#### Polling
Goblox provides a utility method, under [`methodutil`](/packages/methodutil), that can be used to poll a method until it is completed. The [`LuauExecutionTask.TaskInfo()`](/documentation/opencloud/luau-execution#taskinfo) method returns the task information, which we can use to easily reference the task creation information for polling.
```go
package main

import (
    "context"
    "fmt"

    "github.com/typical-developers/goblox/opencloud"
    "github.com/typical-developers/goblox/pkg/methodutil"
)

func main() {
    ctx := context.Background()
    client := opencloud.NewClient().WithAPIKey("YOUR_API_KEY")

    var result int
    var taskError error

    // First, we create the task with the Luau execution API.
    task, _, err := client.LuauExecution.CreateLuauExecutionSessionTask(ctx, "UNIVERSE_ID", "PLACE_ID", nil, opencloud.LuauExecutionTaskCreate{
        Script: opencloud.Pointer("return 1 + 2"),
    })
    if err != nil {
        panic(err)
    }

    // Then, we get the TaskInfo so we can get the task we just created.
    universeId, placeId, versionId, sessionId, taskId := task.TaskInfo()

    // Finally, we poll the task until it's complete and set the reuslt / error in our variables above.
	methodutil.PollMethod(func(done func()) {
		task, resp, err := client.LuauExecution.GetLuauExecutionSessionTask(ctx, universeId, placeId, versionId, sessionId, taskId)
        if err != nil {
            taskError = err
            done()
            return
        }

        // Keeps polling if the ratelimit is exhausted.
        if resp.StatusCode == 429 {
            return
        }

        // Queued means the task is still pending to be executed.
        // Processing means the task is currently being executed.
        // 
        // Only handle the data if the task is done being executed.
        if task.State != opencloud.LuauExecutionStateProcessing && task.State != opencloud.LuauExecutionStateQueued {
            if task.Output != nil && len(task.Output.Results) > 0 {
                result = task.Output.Results[0].(int)
            }

            if task.Error != nil {
                taskError = fmt.Errorf("LuauExecutionTask[%s]: %s", task.Error.Code, task.Error.Message)
            }

            done()
        }
	}, 0)

    fmt.Printf("Result: %d\n", result)
    fmt.Printf("Error: %v\n", taskError)
}
```

Now we can keep checking the status of the task until it is completed and set our result or error into variables.

---

### Using Binary Outputs / Inputs
The Luau execution API also allows you to use binary inputs and outputs. You can either upload a binary input to use in a script or have your script output a binary output. This is useful for large files

You are able to either upload a binary input to use in your sciprt or have your script output the result as a binary output. You must have `EnableBinaryOuput` set to `true` if you want to use a binary output.

#### Uploading a Binary Input
Goblox provides utility methods for uploading a binary input. You can easily use the `UploadLuauExecutionSessionTaskBinaryInput` method to upload a binary input to be used in the task.
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

    var dataBuffer bytes.Buffer
    data := []string{"Hello", ",", "world", "!"} // An example array of strings.
    err := json.newEncoder(&dataBuffer).Encode(data)
    if err != nil {
        panic(err)
    }
    input := dataBuffer.Bytes()

    binaryInput, _, err := client.LuauExecution.CreateLuauExecutionSessionTaskBinaryInput(ctx, "UNIVERSE_ID", opencloud.LuauExecutionSessionTaskBinaryInputCreate{
        Size: opencloud.Pointer(int64(len(input))), // The size of the input you are uploading.
    })
    if err != nil {
        panic(err)
    }
}
```
After we've created the binary input, we can use the `UploadLuauExecutionSessionTaskBinaryInput` method to upload the binary input to `binaryInput.UploadURI`.
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

    var dataBuffer bytes.Buffer
    data := []string{"Hello", ",", "world", "!"} // An example array of strings.
    err := json.newEncoder(&dataBuffer).Encode(data)
    if err != nil {
        panic(err)
    }
    input := dataBuffer.Bytes()

    binaryInput, _, err := client.LuauExecution.CreateLuauExecutionSessionTaskBinaryInput(ctx, "UNIVERSE_ID", opencloud.LuauExecutionSessionTaskBinaryInputCreate{
        Size: opencloud.Pointer(int64(len(input))), // The size of the input you are uploading.
    })
    if err != nil {
        panic(err)
    }

    // Upload the binary input to the URL provided in binaryInput.UploadURI.
    err = client.LuauExecution.UploadLuauExecutionSessionTaskBinaryInput(ctx, binaryInput.UploadURI, input)
    if err != nil {
        panic(err)
    }
}
```
Now that the binary input has been uploaded, it is now available to be used in the script. Below is an example (provided by the OpenCloud API docs) of how to use the binary input in a script.
```luau
local taskInput: LuauExecutionTaskInput = ({...})[1]
local buf: buffer = taskInput.BinaryInput
```

#### Retrieving a Binary Output
You can retrieve a binary output by calling the `BinaryOutput` method on the created task's structure. This method should ONLY be called after the task has been successfully completed.
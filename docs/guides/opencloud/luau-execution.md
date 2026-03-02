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
In order to access the data from the script that was executed, you must poll the endpoint until the task is either complete or fails. 
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
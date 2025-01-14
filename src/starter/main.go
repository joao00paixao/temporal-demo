// starter/main.go
package main

import (
	"context"
	"log"
	"temporal-demo/workflows"

	"go.temporal.io/sdk/client"
)

func main() {
    c, err := client.Dial(client.Options{
        HostPort: "temporal-server:7233",
    })
    if err != nil {
        log.Fatalln("Unable to create client", err)
    }
    defer c.Close()

    workflowOptions := client.StartWorkflowOptions{
        ID:        "sample-workflow",
        TaskQueue: "sample-task-queue",
    }

    we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, workflows.SampleWorkflow, "Temporal")
    if err != nil {
        log.Fatalln("Unable to execute workflow", err)
    }

    var result string
    err = we.Get(context.Background(), &result)
    if err != nil {
        log.Fatalln("Unable to get workflow result", err)
    }
    log.Printf("Workflow result: %v\n", result)
}

// worker/main.go
package main

import (
	"log"
	"temporal-demo/activities"
	"temporal-demo/workflows"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
    c, err := client.Dial(client.Options{
        HostPort: "temporal-server:7233",
    })
    if err != nil {
        log.Fatalln("Unable to create client", err)
    }
    defer c.Close()

    w := worker.New(c, "sample-task-queue", worker.Options{})
    w.RegisterWorkflow(workflows.SampleWorkflow)
    w.RegisterActivity(activities.SampleActivity)

    err = w.Run(worker.InterruptCh())
    if err != nil {
        log.Fatalln("Unable to start worker", err)
    }
}


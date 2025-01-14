// workflows/workflow.go
package workflows

import (
	"temporal-demo/activities"
	"time"

	"go.temporal.io/sdk/workflow"
)

func SampleWorkflow(ctx workflow.Context, name string) (string, error) {
    options := workflow.ActivityOptions{
        StartToCloseTimeout: time.Second * 5,
    }
    ctx = workflow.WithActivityOptions(ctx, options)

    var result string
    err := workflow.ExecuteActivity(ctx, activities.SampleActivity, name).Get(ctx, &result)
    return result, err
}

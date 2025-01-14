// activities/activity.go
package activities

import (
	"context"
	"fmt"
)

func SampleActivity(ctx context.Context, name string) (string, error) {
    return fmt.Sprintf("Hello, %s!", name), nil
}

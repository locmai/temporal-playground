package workflows

import (
	"context"
	"time"
	"fmt"
	"strings"
	"strconv"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"

	"github.com/locmai/temporal-playground/go/shared"
)

// Workflow is a Hello World workflow definition.
func SortWorkflow(ctx workflow.Context, detail shared.SortDetails) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("Sort workflow started", "name", detail.Name)

	var result string
	err := workflow.ExecuteActivity(ctx, SortActivity, detail).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}

	logger.Info("Sort workflow completed.", "result", result)

	return result, nil
}

func SortActivity(ctx context.Context, detail shared.SortDetails) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity", "name", detail.Name)
	var res []int
	switch name := detail.Name; name {
		case "bubble":
			res = bubble(detail.Numbers)
		case "quick":
			res =quick(detail.Numbers, 0, len(detail.Numbers) - 1)
		default:
			fmt.Printf("%s has not been implemented", detail.Name)
		}
	return "[" + joinIntSlice(res,", ") + "]", nil
}


func joinIntSlice(slice []int, separator string) string {
    // Convert each integer to string
    strSlice := make([]string, len(slice))
    for i, num := range slice {
        strSlice[i] = strconv.Itoa(num)
    }
    // Join the string slice using the specified separator
    return strings.Join(strSlice, separator)
}

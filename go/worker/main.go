package main

import (
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/locmai/temporal-playground/go/workflows"
	"github.com/locmai/temporal-playground/go/shared"
)

func main() {
	// The client and worker are heavyweight objects that should be created once per process.
	c, err := client.Dial(client.Options{})
	if err != nil {}
	defer c.Close()

	w := worker.New(c, shared.SORT_TASK_QUEUE_NAME, worker.Options{})

	w.RegisterWorkflow(workflows.SortWorkflow)
	w.RegisterActivity(workflows.SortActivity)

	err = w.Run(worker.InterruptCh())
	if err != nil {

	}
}

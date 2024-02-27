package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"

	"github.com/locmai/temporal-playground/go/workflows"
	"github.com/locmai/temporal-playground/go/shared"
	"github.com/google/uuid"
)

func main() {
	// The client is a heavyweight object that should be created once per process.
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	id := uuid.New()
	workflowOptions := client.StartWorkflowOptions{
		ID:        id.String(),
		TaskQueue: shared.SORT_TASK_QUEUE_NAME,
	}

	sortDetails := shared.SortDetails {
		Name: "quick",
		Numbers: []int{5, 2, 1, 4, 3, 200},
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, workflows.SortWorkflow, sortDetails)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	// Synchronously wait for the workflow completion.
	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", result)
}

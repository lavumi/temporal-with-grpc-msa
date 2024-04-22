package starter

import (
	"context"
	"log"
	etlwork "tempotaletl/internal/temporal"

	"go.temporal.io/sdk/client"
)

func Run() {
	// The client is a heavyweight object that should be created once per process.
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "etl_worlflow",
		TaskQueue: "etl-process",
		// CronSchedule: "0 * * * *",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, etlwork.ETLWorkflow)
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

package etlwork

import (
	"fmt"
	grpcclient "tempotaletl/internal/temporal/grpcClient"
	"time"

	"go.temporal.io/sdk/workflow"
)

func ETLWorkflow(ctx workflow.Context) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("ETL Workflow test start")
	// Extract
	var result grpcclient.Response
	err := workflow.ExecuteActivity(ctx, grpcclient.CallExtractService, 0).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}
	logger.Info("Extract Result", "result", result.Result, " data : ", result.Data)
	// Transform
	err = workflow.ExecuteActivity(ctx, grpcclient.CallTransformService, result.Data).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}
	logger.Info("Transform Result", "result :", result.Result, " data : ", result.Data)
	// Load
	err = workflow.ExecuteActivity(ctx, grpcclient.CallLoadService, result.Data).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}
	logger.Info("Load Result", "result", result.Result, " data : ", result.Data)
	return fmt.Sprint("all process done, result : ", result.Data), nil
}

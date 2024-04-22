package worker

import (
	"log"
	etlwork "tempotaletl/internal/temporal"
	grpcclient "tempotaletl/internal/temporal/grpcClient"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func Run() {
	c, err := client.NewLazyClient(client.Options{
		HostPort: "172.30.1.44:7233",
	})
	if err != nil {
		log.Fatalln("Error creating Temporal client", err)
	}
	defer c.Close()

	workerOptions := worker.Options{}
	w := worker.New(c, "etl-process", workerOptions)
	w.RegisterWorkflow(etlwork.ETLWorkflow)
	w.RegisterActivity(grpcclient.CallExtractService)
	w.RegisterActivity(grpcclient.CallTransformService)
	w.RegisterActivity(grpcclient.CallLoadService)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Error starting worker", err)
	}
}

package main

import (
	"tempotaletl/internal/temporal/starter"
	"tempotaletl/internal/temporal/worker"
)

func main() {

	go worker.Run()
	starter.Run()
}

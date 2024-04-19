package main

import "tempotaletl/internal/msaserver"

func main() {
	msaserver.RunCrawl()
}

//env GOOS=linux go build -o build/crawl-server cmd/crawl.go
//env GOOS=linux go build -o build/extract-server cmd/extract.go
//env GOOS=linux go build -o build/load-server cmd/load.go
//env GOOS=linux go build -o build/transform-server cmd/transform.go

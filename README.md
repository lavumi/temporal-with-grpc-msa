1. Run Temporal with docker  
[Link](https://github.com/temporalio/docker-compose)

1. Install Protobuf
    ```bash
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
    ```
1. Update PATH
    ```bash
    export PATH="$PATH:$(go env GOPATH)/bin"
    ```
1. Build protobuf
    ```bash
    rm -rf api/proto/**
    protoc --proto_path=proto --go_out=api --go_opt=paths=import --go-grpc_out=api --go-grpc_opt=paths=import proto/*.proto
    ```


1. Build
    ```bash
    env GOOS=linux go build -o build/crawl-server cmd/crawl.go
    env GOOS=linux go build -o build/extract-server cmd/extract.go
    env GOOS=linux go build -o build/load-server cmd/load.go
    env GOOS=linux go build -o build/transform-server cmd/transform.go
    env GOOS=linux go build -o build/worker-server cmd/worker.go
    ```

1. Docker image
    ```bash
    docker build -f deployments/crawl.Dockerfile -t crawl-server:0.0.1 .
    docker build -f deployments/extract.Dockerfile -t extract-server:0.0.1 .
    docker build -f deployments/load.Dockerfile -t load-server:0.0.1 .
    docker build -f deployments/transform.Dockerfile -t transform-server:0.0.1 .
    docker build -f deployments/worker.Dockerfile -t worker-server:0.0.1 .
    docker compose up -d
    ```


1. Temporal workflow run
    ```bash
    go run cmd/starter.go
    ```

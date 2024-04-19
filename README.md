

1. Install Protobuf
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
2. update PATH
```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```
3. build protobuf
```bash
rm -rf api/proto/**
protoc --proto_path=proto --go_out=api --go_opt=paths=import --go-grpc_out=api --go-grpc_opt=paths=import proto/*.proto
```


4. build
```bash
env GOOS=linux go build -o build/crawl-server cmd/crawl.go
env GOOS=linux go build -o build/extract-server cmd/extract.go
env GOOS=linux go build -o build/load-server cmd/load.go
env GOOS=linux go build -o build/transform-server cmd/transform.go
```

5. docker image
```bash
docker build -f deployments/crawl.Dockerfile -t crawl-server:0.0.1 .
docker build -f deployments/extract.Dockerfile -t extract-server:0.0.1 .
docker build -f deployments/load.Dockerfile -t load-server:0.0.1 .
docker build -f deployments/transform.Dockerfile -t transform-server:0.0.1 .
docker compose up -d
```


6. temporal run

```bash
go run cmd/temporal.go
```
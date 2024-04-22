docker compose kill
rm -rf api/proto/**
rm -rf build/**
protoc --proto_path=proto --go_out=api --go_opt=paths=import --go-grpc_out=api --go-grpc_opt=paths=import proto/*.proto
env GOOS=linux go build -o build/crawl-server cmd/crawl.go
env GOOS=linux go build -o build/extract-server cmd/extract.go
env GOOS=linux go build -o build/load-server cmd/load.go
env GOOS=linux go build -o build/transform-server cmd/transform.go
env GOOS=linux go build -o build/worker-server cmd/worker.go
docker build -f deployments/crawl.Dockerfile -t crawl-server:0.0.1 .
docker build -f deployments/extract.Dockerfile -t extract-server:0.0.1 .
docker build -f deployments/load.Dockerfile -t load-server:0.0.1 .
docker build -f deployments/transform.Dockerfile -t transform-server:0.0.1 .
docker build -f deployments/worker.Dockerfile -t worker-server:0.0.1 .
docker compose up -d

# docker image prune -f -a
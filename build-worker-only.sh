docker compose kill worker
env GOOS=linux go build -o build/worker-server cmd/worker.go
docker build -f deployments/worker.Dockerfile -t worker-server:0.0.1 .
docker compose up -d worker
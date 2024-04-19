FROM golang:1.22-alpine AS builder

WORKDIR ..

# Go 모듈 파일 복사
COPY go.mod ./
COPY go.sum ./

RUN apk add --no-cache bash
RUN go mod download

COPY api/ ./api/
COPY cmd/ ./cmd/
COPY internal/ ./internal/


RUN CGO_ENABLED=0 GOOS=linux go build -o crawl-server cmd/crawl.go
RUN CGO_ENABLED=0 GOOS=linux go build -o extract-server cmd/extract.go
RUN CGO_ENABLED=0 GOOS=linux go build -o load-server cmd/load.go
RUN CGO_ENABLED=0 GOOS=linux go build -o transform-server cmd/transform.go


FROM alpine AS crawl-image
WORKDIR /root/
COPY --from=builder crawl-server .
CMD ["./crawl-server"]


FROM alpine AS extract-image
WORKDIR /root/
COPY --from=builder extract-server .
CMD ["./extract-server"]


FROM alpine AS load-image
WORKDIR /root/
COPY --from=builder load-server .
CMD ["./load-server"]


FROM alpine AS transform-image
WORKDIR /root/
COPY --from=builder transform-server .
CMD ["./transform-server"]
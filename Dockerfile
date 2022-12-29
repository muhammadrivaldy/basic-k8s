FROM golang:1.19-alpine AS builder

WORKDIR /src/basic_k8s
COPY . .

ENV GO111MODULE=on

RUN go mod tidy && go mod vendor
RUN go build -o main

FROM alpine:latest

COPY --from=builder /src/basic_k8s /src/basic_k8s

WORKDIR /src/basic_k8s

EXPOSE 8080

CMD ["./main"]
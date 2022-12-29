FROM golang:1.17.1-alpine as builder

WORKDIR /app
COPY ../../go.mod ./
COPY ../../go.sum ./
RUN go mod download
COPY ../../  ./

RUN go build main.go

FROM alpine
COPY --from=builder /app /app
ENTRYPOINT ["/app/main"]

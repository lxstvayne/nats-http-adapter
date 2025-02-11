FROM golang:1.23.3 as builder
ARG CGO_ENABLED=0
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o app ./cmd/app/main.go

FROM scratch
COPY --from=builder /app/app /app
ENTRYPOINT ["/app"]

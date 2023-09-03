FROM golang:1.21-bookworm

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go test ./...
RUN go build -o campaign_app ./cmd

CMD ["/app/campaign_app"]

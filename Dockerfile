FROM golang:1.19.1 as builder

WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -ldflags "-s -w" -trimpath -o /app/app ./cmd/app

FROM gcr.io/distroless/static

COPY --from=builder /app/app /home

WORKDIR /home

ENTRYPOINT ["./app"]

EXPOSE 8327

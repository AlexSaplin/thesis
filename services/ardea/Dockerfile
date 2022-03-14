FROM golang:1.14

WORKDIR /build

COPY go.mod go.mod

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /build/app .

COPY config-docker.json .

COPY entrypoint.sh .

ENTRYPOINT ["./entrypoint.sh"]

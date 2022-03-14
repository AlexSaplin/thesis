FROM golang:1.14

WORKDIR /build

RUN go get -v github.com/rakyll/statik

COPY go.mod go.mod

RUN go mod download -x

COPY . .

RUN statik -src=swagger/v1 -p docs -dest=pkg

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /build/app .

COPY config-docker.json .

COPY entrypoint.sh .

ENTRYPOINT ["./entrypoint.sh"]

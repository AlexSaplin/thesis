FROM golang:1.15

WORKDIR /build

COPY go.mod go.mod

RUN go mod download -x

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/main.go

FROM alpine:latest

ENV KUBE_VERSION="v1.18.0"

RUN apk add --update ca-certificates \
 && apk add --update -t deps curl \
 && curl -L https://storage.googleapis.com/kubernetes-release/release/${KUBE_VERSION}/bin/linux/amd64/kubectl -o /usr/local/bin/kubectl \
 && chmod +x /usr/local/bin/kubectl


RUN apk add \
      python3 \
      py3-pip \
  && pip3 install --upgrade pip \
  && pip3 install \
      awscli \
  && apk del --purge deps \
  && rm -rf /var/cache/apk/*

RUN aws --version

WORKDIR /root/

COPY --from=0 /build/app .

COPY service.yaml .

COPY entrypoint.sh .

ENTRYPOINT ["./entrypoint.sh"]

FROM python:3.8

RUN apt-get update && \
    apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common

# install docker client
RUN mkdir -p /tmp/download \
    && curl -L https://download.docker.com/linux/static/stable/x86_64/docker-17.06.2-ce.tgz | tar -xz -C /tmp/download \
    && mv /tmp/download/docker/docker /usr/local/bin/ \
    && rm -rf /tmp/download


WORKDIR /app

COPY app/requirements.txt .

RUN pip install -r requirements.txt

COPY app /app

ENTRYPOINT ["./entrypoint.sh"]

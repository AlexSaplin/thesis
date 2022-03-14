FROM prom/prometheus

WORKDIR /prometheus

COPY prometheus.yml .

COPY entrypoint.sh .

ENTRYPOINT ["sh", "./entrypoint.sh"]

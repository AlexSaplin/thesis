FROM python:3.7

WORKDIR /app

COPY app/requirements.txt .

RUN pip install -r requirements.txt

COPY app /app

ENTRYPOINT ["./entrypoint.sh"]

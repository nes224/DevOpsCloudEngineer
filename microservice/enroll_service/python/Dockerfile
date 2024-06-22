FROM python:3.8-slim
RUN apt-get update && apt-get install -y \
    build-essential \
    pkg-config \
    default-libmysqlclient-dev \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY rpc.py .
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
CMD nameko run rpc --broker amqp://guest:guest@rabbitmq:5672
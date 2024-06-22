FROM python:3.8-slim
WORKDIR /app
COPY api.py .
COPY requirements.txt .
RUN pip install --upgrade pip
RUN pip install --no-cache-dir -r requirements.txt
CMD uvicorn api:app --host 0.0.0.0 --port 7001
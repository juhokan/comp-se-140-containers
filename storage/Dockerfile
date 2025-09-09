FROM python:3.12-slim

WORKDIR /app

RUN pip install flask

COPY storage.py .

EXPOSE 5000

CMD ["python", "storage.py"]

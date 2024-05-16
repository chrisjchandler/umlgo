FROM mirror.gcr.io/library/golang:1.19 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main main.go

FROM mirror.gcr.io/library/python:3.10-slim
WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
COPY . .
COPY --from=builder /app/main .
EXPOSE 8800
CMD ["gunicorn", "-w", "4", "-b", "0.0.0.0:8800", "umlgo:app"]

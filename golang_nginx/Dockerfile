FROM golang:1.22-alpine3.18
WORKDIR /app 
COPY go.mod go.sum ./
RUN go mod download
RUN apk add --update --no-cache bash git openssh
RUN go install github.com/cosmtrek/air@latest
COPY . .
RUN go build -o main . 
EXPOSE 8000

HEALTHCHECK --interval=12s --timeout=12s --start-period=30s --retries=3 \
    CMD curl -fail http://localhost:8000/api/ || exit 1

CMD ["air","./app/main.go","-b","0.0.0.0"]
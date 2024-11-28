FROM golang:1.12.0-alpine3.9

WORKDIR /app

COPY . .

# Install git
RUN apk add --no-cache git

RUN go build -o main .

CMD ["/app/main"]
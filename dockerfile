FROM  golang:1.27-alpine

WORKDIR /app
COPY . . 


RUN go build -o caching-proxy ./cmd/caching-proxy

EXPOSE 8080

CMD ["./caching-proxy", "--port=8080", "--origin=https://dummyjson.com/products"]

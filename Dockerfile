FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o spitter main.go

FROM scratch

COPY --from=builder /app/spitter /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/spitter"]
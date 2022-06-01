FROM golang:alpine AS builder

RUN apk add git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main

FROM chromedp/headless-shell

RUN apt-get update && \
    apt-get install -y dumb-init && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/main .

ENTRYPOINT [ "dumb-init", "--" ]
CMD [ "./main" ]

FROM golang AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main

FROM debian AS runtime

#RUN apt-get update && \
#    apt-get install -y dumb-init musl-dev && \
#    ln -s /usr/lib/x86_64-linux-musl/libc.so /lib/libc.musl-x86_64.so.1 &&\
#    rm -rf /var/lib/apt/lists/*

RUN apt-get update && apt-get install -y chromium

COPY --from=builder /app/main .

#ENTRYPOINT [ "dumb-init", "--" ]
#CMD [ "./main" ]

ENTRYPOINT [ "./main" ]

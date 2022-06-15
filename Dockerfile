FROM golang AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main

#FROM debian AS runtime
#
#RUN apt-get update && apt-get install -y chromium
#
#COPY --from=builder /app/main .
#
#ENTRYPOINT [ "./main" ]

FROM zeke/headless-shell:102.0.5005.115

COPY --from=builder /app/main .

ENTRYPOINT [ "./main" ]

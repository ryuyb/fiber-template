FROM golang:1.24.1 AS builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn make build

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin/ app
COPY --from=builder /src/config.yml data/conf/config.yml

WORKDIR /app

EXPOSE 8080
VOLUME [ "/data/conf" ]

CMD [ "./server", "-conf", "/data/conf" ]

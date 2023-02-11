FROM golang:1.16

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN apt-get update && \
    apt-get install build-essential librdkafka-dev -y

ENV CGO_ENABLED 1

CMD ["tail", "-f" ,"/dev/null"]
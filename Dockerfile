FROM golang:1.11

RUN mkdir -p /tmp
COPY main.go /tmp

CMD ["go", "run", "/tmp/main.go"]

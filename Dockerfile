FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download && \
    go build -o contest-app ./cmd/main.go && \
    chmod +x entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]
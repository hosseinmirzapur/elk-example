FROM golang:1.21-alpine


RUN go mod tidy && \
    go build -o ./bin/elk github.com/hosseinmirzapur/elk-example/cmd/api

EXPOSE 8080 8080

ENTRYPOINT [ "./bin/elk" ]

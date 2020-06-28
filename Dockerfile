FROM golang:1.14-alpine

COPY . /go/src/go-service-template

WORKDIR /go/src/go-service-template

RUN go mod download && go build -o bin/server

EXPOSE ${TEMPLATE_SVC_PORT}

ENTRYPOINT ["bin/server"]

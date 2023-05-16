FROM golang:1.19-alpine3.16

RUN mkdir /app

WORKDIR /app

ADD go.mod .
ADD go.sum .


RUN go mod download
RUN go install -mod=mod github.com/githubnemo/CompileDaemon

ADD . .

EXPOSE 8080

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
FROM public.ecr.aws/docker/library/golang:1.20.2-alpine3.17
WORKDIR /usr/app/src
RUN go install github.com/githubnemo/CompileDaemon@latest
CMD CompileDaemon -command=./main -build="go build -o main ./cmd/main.go" -include=*.go -polling=true
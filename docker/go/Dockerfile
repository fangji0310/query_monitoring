FROM public.ecr.aws/docker/library/golang:1.20.2-alpine3.17
WORKDIR /usr/app/src
RUN go install github.com/githubnemo/CompileDaemon@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
#CMD CompileDaemon -command=./main -build="go build -o main ./cmd/main.go" -include=*.go -polling=true
CMD CompileDaemon -command="dlv --listen=:31222 --headless=true --api-version=2 exec ./main" -build="go build -gcflags=all=-N -gcflags=-l -o main ./cmd/main.go" -pattern="(.+\.go|.+\.c|.+\.yaml)$" -directory="/usr/app/src" -polling=true -graceful-kill=true
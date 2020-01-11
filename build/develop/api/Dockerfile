FROM cisco/base-golang

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

ENTRYPOINT CompileDaemon -build='go build -o ./bin/dev-cisco-api ./cmd/api' -command=./bin/dev-cisco-api
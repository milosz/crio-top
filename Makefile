build:
	go build -o bin/crio-top src/crio-top/main.go

run:
	go run src/crio-top/main.go --configuration examples/configuration.yaml

test:
	go test ./src/...

coverage:
	go test -cover ./src/...

all: build
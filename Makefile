build:
	go build -o bin/crio-top src/crio-top/main.go

run:
	go run src/crio-top/main.go --configuration examples/configuration.yaml

all: build
.PHONY: get_delve startDockerDev start_debug_build build

get_delve:
	go get github.com/go-delve/delve/cmd/dlv

runMain:
	go run main.go

startDockerDev:
	docker-compose up

start_debug_build:
	docker-compose up --build

build:
	go build -o ./build/quickstart main.go

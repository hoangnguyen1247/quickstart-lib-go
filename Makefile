.PHONY: get_delve start_debug start_debug_build

get_delve:
    go get github.com/go-delve/delve/cmd/dlv

start_debug:
    docker-compose up

start_debug_build:
    docker-compose up --build

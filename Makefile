.PHONY: dev
dev: 
	make build && make run

.PHONY: build
build:
	go build .

.PHONY: run
run:
	go run .
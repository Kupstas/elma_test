default: run

.PHONY: test
test:
	go test --race --count=5 ./...

.PHONY: run
run:
	go run main.go
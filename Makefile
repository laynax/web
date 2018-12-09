all:
	go build

run: all
	./web

test:
	go test
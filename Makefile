build:
	go build ./cmd/app/main.go ./bin/app

run:	build
	./bin/app

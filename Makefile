ensure:
	dep ensure

test: ensure
	go test --cover ./...

build: ensure
	go build -o ./bin/gameoflife main.go

run:
	./bin/gameoflife

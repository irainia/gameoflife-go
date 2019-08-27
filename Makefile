ensure:
	dep ensure

test: ensure
	go test --cover ./...

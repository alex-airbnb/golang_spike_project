install:
	go get .

format:
	gofmt -w ./..

lint:
	gofmt -d ./..

test:
	go test -v ./...
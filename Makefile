.PHONY: build test install

build:
	@echo "Building the application..."
	go build -o jwt-parser main.go

test:
	@echo "Running tests..."
	go test -v ./...

install:
	@echo "Installing the application..."
	go install

# Makefile

# Variables
APP_NAME := email-gateway
DIST := dist
DOCKER_IMAGE := $(APP_NAME):latest

# Build the Go binary
build:
	go build -o $(DIST)/$(APP_NAME) main.go

# Run the Go application
run: build
	$(DIST)/$(APP_NAME)

# Build the Docker image
docker:
	docker build -t $(DOCKER_IMAGE) .

# Run the Docker container
docker-run:
	docker run --rm -p 5000:5000 -v ./config.json:/app/config.json --name $(APP_NAME) $(DOCKER_IMAGE)

# Clean up the built binary
clean:
	rm -rf $(DIST)


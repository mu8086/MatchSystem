# Makefile for building the Go application in a minimal Docker container

# Variables
APP_NAME := MatchSystem
DOCKER_IMAGE := mu8086/match-system
DOCKER_PORT_MAPPING := -p 8080:8080

# Build the Docker image
docker-build:
	docker build -t $(DOCKER_IMAGE) .

# Run the Docker container
docker-run:
	docker run -d $(DOCKER_PORT_MAPPING) --name $(APP_NAME) $(DOCKER_IMAGE)

# Stop and remove the Docker container
docker-stop:
	docker stop $(APP_NAME) && docker rm $(APP_NAME)

# Default target (build Docker image)
default: docker-build

# Use the official Go Alpine image as the base image for building the application.
# Alpine is chosen for its lightweight nature, making it suitable for containers.
FROM golang:alpine as BuildStage

# Update the package index and install necessary dependencies:
# - `git`: For managing Go modules that might be hosted on Git repositories.
# - `ca-certificates`: To enable HTTPS for external communications.
# - `tzdata`: Provides time zone data for time-related operations.
# - `gcc` and `libc-dev`: Required for compiling C/C++ dependencies, if needed.
# - `openssh-client`: To support SSH-based Git operations.
RUN apk update && \
    apk add --no-cache git ca-certificates tzdata gcc libc-dev openssh-client && \
    update-ca-certificates

# Set the working directory inside the container.
# This is where the application code will reside and be executed.
WORKDIR /go/src/github.com/salehrashid/mini-project/cmd/web_mini_project

# Specify the command to run the Go application.
# `go run main.go` compiles and executes the application directly.
CMD go run main.go

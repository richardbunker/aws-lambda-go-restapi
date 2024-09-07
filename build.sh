#!/bin/bash

# Build the project
echo "Building the project..."
OOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o bootstrap -tags lambda.norpc main.go
echo "Build complete."

# Zip the binary
echo "Zipping the binary..."
zip deployment.zip bootstrap
echo "Zip complete."

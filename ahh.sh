#!/bin/bash

# Create directories
mkdir -p cmd internal pkg build docs tests

# Create subdirectories within internal
mkdir -p internal/handler internal/logger internal/parser internal/response

# Create files
touch cmd/main.go
touch internal/handler/handler.go
touch internal/logger/logger.go
touch internal/parser/rparser.go
touch internal/response/responsewriter.go
touch go.mod
touch go.sum
touch README.md

# Move existing files to their respective directories (optional)
mv main.go cmd/
mv handler.go internal/handler/
mv logger.go internal/logger/
mv rparser.go internal/parser/
mv responsewritter.go internal/response/

echo "Project structure created successfully!"

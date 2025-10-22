#!/bin/bash

# Exit immediately if a command fails
set -e

# Print each command before executing
set -x

# Navigate to the project root (optional, adjust path if needed)
cd "$(dirname "$0")"

# Build the Go binary
echo "🏗️ Building Go API..."
go build ./cmd/api

# Run the binary
echo "🚀 Starting API server..."
./api


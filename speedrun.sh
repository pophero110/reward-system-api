#!/bin/bash

# Exit immediately if a command fails
set -e
# Print each command before executing
set -x

# Navigate to the project root (adjust path if needed)
cd "$(dirname "$0")"

# Function to build and run the API
run_server() {
  echo "🏗️ Building Go API..."
  go build -o api ./cmd/api

  echo "🚀 Starting API server..."
  ./api
}

# Trap Ctrl+C to not exit the script
trap 'echo "🛑 Server stopped. Use 'r' to restart or Ctrl+D to quit."' SIGINT

# Interactive loop
while true; do
  echo ""
  echo "Options: [r] restart server, [q] quit"
  read -p "> " choice || break  # Ctrl+D exits loop
  case "$choice" in
    r|R)
      # Kill any running API process (if any)
      pkill -f "./api" || true
      run_server &
      ;;
    q|Q)
      echo "👋 Exiting..."
      pkill -f "./api" || true
      break
      ;;
    *)
      echo "⚠️ Invalid option: $choice"
      ;;
  esac
done

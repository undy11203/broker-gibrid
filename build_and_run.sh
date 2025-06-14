#!/bin/bash

set -e

echo "Commands available:"
echo "  build-core    - Build core C++ broker"
echo "  build-server  - Build Go server"
echo "  run-server    - Run Go server (which launches core broker)"
echo "  all           - Build core, build server, and run server"

case "$1" in
  build-core)
    echo "Building core C++ broker..."
    mkdir -p build
    cd build
    cmake ../core
    cmake --build .
    cd ..
    ;;
  build-server)
    echo "Building Go server..."
    go build -o server main.go
    ;;
  run-server)
    echo "Running Go server..."
    ./server
    ;;
  all)
    echo "Building core C++ broker..."
    mkdir -p build
    cd build
    cmake ../core
    cmake --build .
    cd ..
    echo "Building Go server..."
    go build -o server main.go
    echo "Running Go server..."
    ./server
    ;;
  *)
    echo "Usage: $0 {build-core|build-server|run-server|all}"
    exit 1
    ;;
esac

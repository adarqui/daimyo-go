#!/bin/bash

while true; do
    inotifywait -r src --exclude "swp" -e modify
    go test ./...
done

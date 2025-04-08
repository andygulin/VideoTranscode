#!/bin/bash

GO_PROGRAM="main.go"
GO_OUTPUT="VideoTranscode"

if [ ! -f "$GO_PROGRAM" ]; then
    echo "Error: $GO_PROGRAM not found."
    exit 1
fi

PLATFORMS=("darwin amd64 " "windows amd64 .exe")

for platform in "${PLATFORMS[@]}"; do
    IFS=' ' read -r -a parts <<< "$platform"
    GOOS=${parts[0]}
    GOARCH=${parts[1]}
    EXTENSION=${parts[2]}

    echo "Compiling $GO_PROGRAM for $GOOS $GOARCH..."
    (
        export GOOS GOARCH
        go build -o "${GO_OUTPUT}_${GOOS}_${GOARCH}${EXTENSION}" $GO_PROGRAM
    )
    if [ $? -ne 0 ]; then
        echo "Error: Failed to compile for $GOOS $GOARCH."
        exit 1
    fi
done

echo "Compilation completed successfully."
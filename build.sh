#!/bin/bash

# 64 but Intel Windows
env GOOS=windows GOARCH=amd64 go build -o bin/Win64/testpdfgen.exe testpdfgen.go

# # 64 bit Intel Mac
env GOOS=darwin GOARCH=amd64 go build -o bin/OSX-Intel/testpdfgen testpdfgen.go

# # Apple Silicon
env GOOS=darwin env GOARCH=arm64 go build -o bin/OSX-AppleSilicon/testpdfgen testpdfgen.go
env GOOS=darwin env GOARCH=arm64 go build -o testpdfgen testpdfgen.go


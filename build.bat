@echo off
setlocal

set BINARY_NAME=myapp
set VERSION=1.0.0

echo Building for Windows...
set GOOS=windows&& set GOARCH=amd64&& go build -ldflags "-s -w" -o dist/windows-amd64.exe

echo Building for Linux...
set GOOS=linux&& set GOARCH=amd64&& go build -ldflags "-s -w" -o dist/linux-amd64

echo Building for macOS ARM64...
set GOOS=darwin&& set GOARCH=arm64&& go build -ldflags "-s -w" -o dist/osx-arm64

echo Build completed!
endlocal

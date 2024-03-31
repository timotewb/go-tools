# go-tools-cli

### Build commands
#### Mac
GOOS=darwin GOARCH=arm64 go build -o build/go-tools-cli
#### Raspberry PI
env GOOS=linux GOARCH=arm GOARM=7 go build -o build/go-tools-cli

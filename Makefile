APP_NAME=yon
DIST=dist
GO=go

LDFLAGS=-ldflags="-s -w"

build:
	@mkdir -p $(DIST)

	GOOS=darwin GOARCH=arm64 $(GO) build $(LDFLAGS) -o $(DIST)/$(APP_NAME)-darwin-arm64
	GOOS=darwin GOARCH=amd64 $(GO) build $(LDFLAGS) -o $(DIST)/$(APP_NAME)-darwin-amd64
	GOOS=linux GOARCH=amd64 $(GO) build $(LDFLAGS) -o $(DIST)/$(APP_NAME)-linux-amd64
	GOOS=linux GOARCH=arm64 $(GO) build $(LDFLAGS) -o $(DIST)/$(APP_NAME)-linux-arm64
	GOOS=windows GOARCH=amd64 $(GO) build $(LDFLAGS) -o $(DIST)/$(APP_NAME)-windows-amd64.exe
	GOOS=windows GOARCH=arm64 $(GO) build $(LDFLAGS) -o $(DIST)/$(APP_NAME)-windows-arm64.exe


checksum:
	cd $(DIST) && shasum -a 256 * > checksums.txt

clean:
	rm -rf $(DIST)
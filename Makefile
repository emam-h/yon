APP_NAME=yon
DIST=dist
GO=go

LDFLAGS=-ldflags="-s -w"

build:
	@mkdir -p $(DIST)

	GOOS=darwin GOARCH=arm64 $(GO) build $(LDFLAGS) -o $(DIST)/yon-darwin-arm64
	GOOS=darwin GOARCH=amd64 $(GO) build $(LDFLAGS) -o $(DIST)/yon-darwin-amd64
	GOOS=linux GOARCH=amd64 $(GO) build $(LDFLAGS) -o $(DIST)/yon-linux-amd64
	GOOS=linux GOARCH=arm64 $(GO) build $(LDFLAGS) -o $(DIST)/yon-linux-arm64
	GOOS=windows GOARCH=amd64 $(GO) build $(LDFLAGS) -o $(DIST)/yon-windows-amd64.exe
	GOOS=windows GOARCH=arm64 $(GO) build $(LDFLAGS) -o $(DIST)/yon-windows-arm64.exe

checksum:
	cd $(DIST) && shasum -a 256 * > checksums.txt

package:
	@echo "Packaging for Homebrew..."
	@cd $(DIST) && \
	cp yon-darwin-arm64 yon && tar -czf yon-darwin-arm64.tar.gz yon && rm yon && \
	cp yon-darwin-amd64 yon && tar -czf yon-darwin-amd64.tar.gz yon && rm yon && \
	cp yon-linux-amd64 yon && tar -czf yon-linux-amd64.tar.gz yon && rm yon && \
	cp yon-linux-arm64 yon && tar -czf yon-linux-arm64.tar.gz yon && rm yon

clean:
	rm -rf $(DIST)
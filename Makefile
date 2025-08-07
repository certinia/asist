include .env

PACKAGE_NAME ?= asist.vsix
PREFIX=github.com/certinia/asist/scanner
LDFLAGS=-s -w -X '$(PREFIX).Version=$(VERSION)'

build:
	go build -ldflags "$(LDFLAGS)" .

install:
	go install -ldflags "$(LDFLAGS)" .

create-rule-mapping:
	cd cmd/gen-models && go run main.go

integration-tests:
	go test github.com/certinia/asist/integrationtests

local-test:
	go run . -c .asist.example.yaml integrationtests

unit-test:
	go test ./... -cover

build-binaries:
	GOOS=windows GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o asist.exe
	GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o asist_darwin_amd64
	GOOS=darwin GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o asist_darwin_arm64
	GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o asist_linux_amd64
	GOOS=linux GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o asist_linux_arm64

build-binaries-for-extension: build-binaries
	cp ./asist.exe ./extension/asist.exe
	cp ./asist_darwin_amd64 ./extension/asist_darwin_amd64
	cp ./asist_darwin_arm64 ./extension/asist_darwin_arm64
	cp ./asist_linux_amd64 ./extension/asist_linux_amd64
	cp ./asist_linux_arm64 ./extension/asist_linux_arm64

copy-config-file:
	cp .asist.example.yaml ./extension/client/configfiletemplate.yaml
	cp LICENSE.txt ./extension/LICENSE.txt
	cp changelog.md ./extension/changelog.md

build-vscode-extension: copy-config-file build-binaries-for-extension
	cd extension && npm install && npm run compile && vsce package -o '${PACKAGE_NAME}'

build-vscode-extension-prepublish: copy-config-file build-binaries-for-extension
	cd extension && npm install && npm run prepublishOnly

build-vscode-extension-binary-exist: copy-config-file
	cd extension && npm install && vsce package -o '${PACKAGE_NAME}'

build-linux:
	cd acli && GOOS=linux GOARCH=amd64 go build -ldflags="-w -s"

build-windows:
	cd acli && GOOS=windows GOARCH=amd64 go build -ldflags="-w -s"

build-all: build-linux build-windows
	@echo ""

deploy: build-linux
	mv acli/acli ~/.local/bin/acli
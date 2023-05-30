build:
	cd acli && GOOS=linux GOARCH=amd64 go build -ldflags="-w -s"

deploy: build
	mv acli/acli ~/.local/bin/acli
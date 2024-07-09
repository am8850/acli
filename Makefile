# Get version from git hash
git_hash := $(shell git rev-parse --short HEAD || echo 'development')

# Get current date
current_time = $(shell date +"%Y-%m-%d:T%H:%M:%S")

# Add linker flags
linker_flags = '-s -w -X main.buildTime=${current_time} -X main.version=0.1.${git_hash}'

build-linux:
	@echo ${linker_flags}
	cd acli && GOOS=linux GOARCH=amd64 go build -ldflags=${linker_flags} -o acli
	cd acli && tar -czvf acli.linux.amd64.tar.gz acli

build-windows:
	cd acli && GOOS=windows GOARCH=amd64 go build -ldflags=${linker_flags} -o acli.exe
	cd acli && zip acli.windows.amd64.zip acli.exe

build-mac:
	cd acli && GOOS=darwin GOARCH=amd64 go build -ldflags=${linker_flags} -o acli.mac.amd64
	cd acli && tar -czvf acli.mac.amd64.tar.gz acli
	cd acli && GOOS=darwin GOARCH=arm64 go build -ldflags=${linker_flags} -o acli.mac.arm64
	cd acli && tar -czvf acli.mac.arm64.tar.gz acli

build-all: build-linux build-windows build-mac
	@echo ""

install:
	cd acli && GOOS=linux GOARCH=amd64 go build -ldflags=${linker_flags} -o acli
	cd acli && sudo mv acli /usr/local/bin
	@echo "Published to /usr/local/bin/acli"


deploy: build-linux
	mv acli/acli ~/.local/bin/acli

clean:
	-rm acli/acli
	-rm acli/acli.exe
	-rm acli/acli.linux.amd64.tar.gz 
	-rm acli/acli.windows.amd64.zip
	-rm acli/acli.mac.amd64
	-rm acli/acli.mac.amd64.tar.gz
	-rm acli/acli.mac.arm64
	-rm acli/acli.mac.arm64.tar.gz

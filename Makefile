build:
	cd acli && go build .

deploy: build
	mv acli/acli ~/.local/bin/acli
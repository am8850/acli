default:
	@echo ""
	
buildui:
	@echo "Building the UI..."
	cd frontend && npm run build

prod:
	@echo "Building the server..."
	cd server && GOOS=linux GOARCH=amd64 go build -ldflags="-w -s"

runui:
	@echo "Running from the UI..."
	cd frontend && npm run dev

run: buildui
	@echo "Running from the server..."
	rm -rf server/public
	mkdir -p server/public
    cp -r frontend/dist/. server/public/.
	cd server && go run .

post-install:	
	cp -r post-frontend/. frontend/.
	rm -rf post-frontend/

install-git:
	git init
	git branch -m main
	npm create vite@latest frontend -- --template react-ts
	cd frontend && npm install && npm install axios && npm install -D tailwindcss postcss autoprefixer && npx tailwindcss init -p
	cd server && go mod tidy
	make post-install
	

install:
	npm create vite@latest frontend -- --template react-ts
	cd frontend && npm install && npm install axios && npm install -D tailwindcss postcss autoprefixer && npx tailwindcss init -p
	make post-install

infra:
	cd infrastructure && terraform init && terraform apply

docker:
	cd server && docker build . -t server:latest
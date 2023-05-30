nada:
	@echo ""
	
buildui:
	cd frontend && npm run build

runui:
	cd frontend && npm run dev

run: buildui
	cd server && go run .

post-install:
	cp -r post-frontend/. frontend/.
	rm -rf post-frontend/

install:
	npm create vite@latest frontend -- --template react-ts
	cd frontend && npm install && npm install axios && npm install -D tailwindcss postcss autoprefixer && npx tailwindcss init -p
	make post-install

infra:
	cd infrastructure && terraform init && terraform apply

docker:
	cd server && docker build . -t server:latest
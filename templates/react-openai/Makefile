default:
	@echo ""

buildui:
	@echo "Building the UI..."
	cd frontend && npm run build

prod: buildui
	@echo ""

runui:
	@echo "Running from the UI..."
	cd frontend && npm run dev

run: runui
	@echo ""

post-install:
	@echo "Post frontend install..."
	cp -r post-frontend/. frontend/.
	rm -rf post-frontend

init-git:
	@echo "Initializing git..."
	git init
	git branch -m main

frontend-install:
	@echo "frontend install..."
	npm create vite@latest frontend -- --template react-ts
	cd frontend && npm install && npm install axios && npm install -D tailwindcss postcss autoprefixer && npx tailwindcss init -p
	cd frontend && npm i react-router-dom && npm i @headlessui/react && npm i @heroicons/react && npm i @tailwindcss/forms
	cd frontend && npm i flowbite-react && npm i react-markdown && npm i remark-gfm

server-install:
	@echo ""

git-install: init-git frontend-install post-install
	@echo ""

install: frontend-install post-install
	@echo ""

infra:
	cd infrastructure && terraform init && terraform apply

docker:
	@echo "You may want to rename the image to something more meaningful and with versions"

podman:
	@echo "You may want to rename the image to something more meaningful and with versions"

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
	cd frontend && npm i react-router-dom && npm i react-markdown && npm i remark-gfm
	@echo "Make sure to add an .env file on the frontend folder, and add the following values from Azure."
	@echo "VITE_OPENAI_KEY=<THE API KEY""
	@echo "VITE_OPENAI_DAVINCI_URL=<FULL URL WITH THE OPENAI MODEL AND VERSION>"
	@echo "VITE_OPENAI_GTP_UR=<FULL URL WITH THE OPENAI MODEL AND VERSION>"

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

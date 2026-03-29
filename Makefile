up-d:
	docker compose up -d

lint:
	golangci-lint run

tidy:
	go mod tidy

build-ui:
	cd ui && npm run build-prod

build-image: build-ui
	docker build --target final -t rest-template .

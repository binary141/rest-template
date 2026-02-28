up-d:
	docker compose up -d

lint:
	golangci-lint run

tidy:
	go mod tidy

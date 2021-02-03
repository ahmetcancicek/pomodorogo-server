postgres:
	docker-compose up -d postgres
run:
	go run cmd/server/main.go
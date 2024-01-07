dev:
	docker-compose -f docker-compose.dev.yml down
	docker-compose -f docker-compose.dev.yml up -d --build

dev-logs:
	docker-compose -f docker-compose.dev.yml logs

dev-down:
	docker-compose -f docker-compose.dev.yml down

create-migration:
	goose -dir database/migrations create $(FILENAME) sql

migrate:
	go run . migrate
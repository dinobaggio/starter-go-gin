dev:
	docker-compose -f docker-compose.dev.yml down
	docker-compose -f docker-compose.dev.yml up -d

dev-logs:
	docker-compose -f docker-compose.dev.yml logs

dev-down:
	docker-compose -f docker-compose.dev.yml down

dev-build:
	docker-compose -f docker-compose.dev.yml build

create-migration:
	goose -dir database/migrations create $(FILENAME) sql
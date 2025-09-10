# Variabili configurabili
DB_CONTAINER_NAME = amilate-db
DB_USER = user
DB_PASS = pass
DB_NAME = amilate
DB_PORT = 5432
MIGRATIONS_DIR = migrations
DB_URL = postgres://$(DB_USER):$(DB_PASS)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable

# Default target
.PHONY: run
run: db-up wait-for-db migrate-up app-run

# 1. Avvia il container Postgres (se non esiste già)
.PHONY: db-up
db-up:
	@if [ ! "$$(docker ps -q -f name=$(DB_CONTAINER_NAME))" ]; then \
		if [ "$$(docker ps -aq -f status=exited -f name=$(DB_CONTAINER_NAME))" ]; then \
			echo "🔄 Starting existing Postgres container..."; \
			docker start $(DB_CONTAINER_NAME); \
		else \
			echo "🐘 Creating new Postgres container..."; \
			docker run --name $(DB_CONTAINER_NAME) \
				-e POSTGRES_USER=$(DB_USER) \
				-e POSTGRES_PASSWORD=$(DB_PASS) \
				-e POSTGRES_DB=$(DB_NAME) \
				-p $(DB_PORT):5432 \
				-d postgres:15; \
		fi \
	else \
		echo "✅ Postgres container already running"; \
	fi

.PHONY: wait-for-db
wait-for-db:
	@echo "⏳ Waiting for Postgres to be ready..."
	@until docker exec $(DB_CONTAINER_NAME) pg_isready -U $(DB_USER) > /dev/null 2>&1; do \
		sleep 1; \
	done
	@echo "✅ Postgres is ready!"

.PHONY: migrate-up
migrate-up:
	@echo "🚀 Running migrations..."
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

.PHONY: migrate-down
migrate-down:
	@echo "⏪ Rolling back last migration..."
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down 1

.PHONY: app-run
app-run:
	@echo "▶️ Starting application..."
	go run ./cmd/amilate/main.go

.PHONY: db-reset
db-reset:
	@echo "🔥 Resetting database..."
	docker exec -i $(DB_CONTAINER_NAME) dropdb -U $(DB_USER) --if-exists $(DB_NAME)
	docker exec -i $(DB_CONTAINER_NAME) createdb -U $(DB_USER) $(DB_NAME)
	make migrate-up

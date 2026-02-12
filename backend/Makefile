# Variabili configurabili
DB_CONTAINER_NAME = amilate-db
DB_USER = user
DB_PASS = pass
DB_NAME = amilate
DB_PORT = 5432
MIGRATIONS_DIR = migrations
DB_URL = postgres://$(DB_USER):$(DB_PASS)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable

# Test variables
TEST_COVERAGE_DIR = coverage
TEST_COVERAGE_FILE = $(TEST_COVERAGE_DIR)/coverage.out
TEST_COVERAGE_HTML = $(TEST_COVERAGE_DIR)/coverage.html

# Default target
.PHONY: run
run: db-up wait-for-db migrate-up app-run

# ===== DATABASE COMMANDS =====

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
		echo "Postgres container already running"; \
	fi

.PHONY: db-stop
db-stop:
	@echo "Stopping Postgres container..."
	docker stop $(DB_CONTAINER_NAME)

.PHONY: db-remove
db-remove: db-stop
	@echo "Removing Postgres container..."
	docker rm $(DB_CONTAINER_NAME)

.PHONY: wait-for-db
wait-for-db:
	@echo "Waiting for Postgres to be ready..."
	@until docker exec $(DB_CONTAINER_NAME) pg_isready -U $(DB_USER) > /dev/null 2>&1; do \
		sleep 1; \
	done
	@echo "Postgres is ready!"

.PHONY: db-reset
db-reset:
	@echo "Resetting database..."
	docker exec -i $(DB_CONTAINER_NAME) dropdb -U $(DB_USER) --if-exists $(DB_NAME)
	docker exec -i $(DB_CONTAINER_NAME) createdb -U $(DB_USER) $(DB_NAME)
	make migrate-up

# ===== MIGRATION COMMANDS =====

.PHONY: migrate-up
migrate-up:
	@echo "Running migrations..."
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

.PHONY: migrate-down
migrate-down:
	@echo "Rolling back last migration..."
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down 1

.PHONY: migrate-down-all
migrate-down-all:
	@echo "Rolling back all migrations..."
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down

.PHONY: migrate-create
migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $$name

# ===== APPLICATION COMMANDS =====

.PHONY: app-run
app-run:
	@echo "▶️  Starting application..."
	go run ./cmd/amilate/main.go

.PHONY: app-build
app-build:
	@echo "Building application..."
	go build -o bin/amilate ./cmd/amilate/main.go

.PHONY: app-clean
app-clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/

# ===== MOCK GENERATION =====

.PHONY: mocks
mocks:
	@echo "Generating mocks..."
	@mkdir -p test/mocks
	mockery --name=UserRepositoryInterface --dir=./internal/repository/interface --output=./test/mocks --outpkg=mocks
	mockery --name=AppointmentRepositoryInterface --dir=./internal/repository/interface --output=./test/mocks --outpkg=mocks
	mockery --name=DestinationRepositoryInterface --dir=./internal/repository/interface --output=./test/mocks --outpkg=mocks
	mockery --name=UserServiceInterface --dir=./internal/service/interface --output=./test/mocks --outpkg=mocks
	mockery --name=AuthServiceInterface --dir=./internal/service/interface --output=./test/mocks --outpkg=mocks
	mockery --name=AppointmentServiceInterface --dir=./internal/service/interface --output=./test/mocks --outpkg=mocks
	mockery --name=DestinationServiceInterface --dir=./internal/service/interface --output=./test/mocks --outpkg=mocks
	@echo "Mocks generated in test/mocks/"

.PHONY: mocks-clean
mocks-clean:
	@echo "Cleaning mocks..."
	rm -rf test/mocks/*

# ===== TEST COMMANDS =====

.PHONY: test
test:
	@echo "Running all tests..."
	go test -v ./test/...

.PHONY: test-unit
test-unit:
	@echo "Running unit tests..."
	go test -v ./test/*_test.go

.PHONY: test-service
test-service:
	@echo "Running service tests..."
	go test -v ./test -run "Service"

.PHONY: test-handler
test-handler:
	@echo "Running handler tests..."
	go test -v ./test -run "Handler"

.PHONY: test-coverage
test-coverage:
	@echo "Running tests with coverage..."
	@mkdir -p $(TEST_COVERAGE_DIR)
	go test -v -coverprofile=$(TEST_COVERAGE_FILE) ./test/...
	go tool cover -html=$(TEST_COVERAGE_FILE) -o $(TEST_COVERAGE_HTML)
	@echo "Coverage report generated: $(TEST_COVERAGE_HTML)"

.PHONY: test-coverage-open
test-coverage-open: test-coverage
	@echo "🌐 Opening coverage report in browser..."
	@open $(TEST_COVERAGE_HTML) || xdg-open $(TEST_COVERAGE_HTML)

.PHONY: test-watch
test-watch:
	@echo "👀 Watching tests..."
	@which nodemon > /dev/null || (echo "nodemon not installed. Run: npm install -g nodemon" && exit 1)
	nodemon --exec "make test" --watch internal --watch test --ext go

# ===== CODE QUALITY =====

.PHONY: lint
lint:
	@echo "Running linter..."
	golangci-lint run ./...

.PHONY: fmt
fmt:
	@echo "Formatting code..."
	go fmt ./...

.PHONY: vet
vet:
	@echo "Running go vet..."
	go vet ./...

.PHONY: check
check: fmt vet lint test
	@echo "All checks passed!"

# ===== DEPENDENCIES =====

.PHONY: deps
deps:
	@echo "Downloading dependencies..."
	go mod download

.PHONY: deps-tidy
deps-tidy:
	@echo "Tidying dependencies..."
	go mod tidy

.PHONY: deps-update
deps-update:
	@echo "Updating dependencies..."
	go get -u ./...
	go mod tidy

# ===== DOCKER COMMANDS =====

.PHONY: docker-build
docker-build:
	@echo "Building Docker image..."
	docker build -t amilate:latest .

.PHONY: docker-run
docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 amilate:latest

# ===== UTILITIES =====

.PHONY: clean
clean: app-clean mocks-clean
	@echo "Cleaning all artifacts..."
	rm -rf $(TEST_COVERAGE_DIR)

.PHONY: help
help:
	@echo "Available commands:"
	@echo ""
	@echo "Database:"
	@echo "  make db-up          - Start Postgres container"
	@echo "  make db-stop        - Stop Postgres container"
	@echo "  make db-remove      - Remove Postgres container"
	@echo "  make db-reset       - Reset database (drop + create + migrate)"
	@echo ""
	@echo "Migrations:"
	@echo "  make migrate-up     - Run all migrations"
	@echo "  make migrate-down   - Rollback last migration"
	@echo "  make migrate-down-all - Rollback all migrations"
	@echo "  make migrate-create - Create new migration"
	@echo ""
	@echo "Application:"
	@echo "  make run            - Start DB, run migrations, start app"
	@echo "  make app-run        - Start application only"
	@echo "  make app-build      - Build application binary"
	@echo ""
	@echo "Mocks:"
	@echo "  make mocks          - Generate all mocks"
	@echo "  make mocks-clean    - Remove all mocks"
	@echo ""
	@echo "Tests:"
	@echo "  make test           - Run all tests"
	@echo "  make test-unit      - Run unit tests only"
	@echo "  make test-service   - Run service tests only"
	@echo "  make test-handler   - Run handler tests only"
	@echo "  make test-coverage  - Run tests with coverage report"
	@echo "  make test-coverage-open - Run tests and open coverage in browser"
	@echo "  make test-watch     - Watch and run tests on file changes"
	@echo ""
	@echo "Code Quality:"
	@echo "  make lint           - Run linter"
	@echo "  make fmt            - Format code"
	@echo "  make vet            - Run go vet"
	@echo "  make check          - Run all checks (fmt + vet + lint + test)"
	@echo ""
	@echo "Dependencies:"
	@echo "  make deps           - Download dependencies"
	@echo "  make deps-tidy      - Tidy dependencies"
	@echo "  make deps-update    - Update all dependencies"
	@echo ""
	@echo "Cleanup:"
	@echo "  make clean          - Clean all artifacts"

.PHONY: all
all: deps mocks check app-build
	@echo "Build complete!"
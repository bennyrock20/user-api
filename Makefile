# Makefile for Docker Compose with Go and PostgreSQL

# Default variables (can be overridden)
COMPOSE_FILE ?= docker-compose.yml
DOCKER_COMPOSE ?= docker-compose-v1

# Targets

# Default target
.PHONY: up
up: ## Bring up the Docker Compose services
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) up

.PHONY: build
build: ## Bring up the Docker Compose services
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) build

.PHONY: down
down: ## Take down the Docker Compose services
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) down

.PHONY: test
test: ## Run all tests cases
	go test ./...

.PHONY: logs
logs: ## Tail the logs for the Docker Compose services
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) logs -f

.PHONY: restart
restart: ## Restart the Docker Compose services
	$(MAKE) down COMPOSE_FILE=$(COMPOSE_FILE) DOCKER_COMPOSE=$(DOCKER_COMPOSE)
	$(MAKE) up COMPOSE_FILE=$(COMPOSE_FILE) DOCKER_COMPOSE=$(DOCKER_COMPOSE)

.PHONY: clean
clean: ## Clean up all resources
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) down -v --rmi all --remove-orphans

.PHONY: jwt
jwt:
	docker compose run web go run main.go jwt
	
.PHONY: help
help: ## Show this help
	@echo "Usage: make [target] [COMPOSE_FILE=your-compose-file.yml] [DOCKER_COMPOSE=your-compose-command]"
	@echo ""
	@echo "Targets:"
	@awk '/^# Targets/,0' $(MAKEFILE_LIST) | grep -E '^[a-zA-Z_-]+:' | sort | awk -F ':


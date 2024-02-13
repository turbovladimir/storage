define setup_env
    $(eval ENV_FILE := .env.$(1))
    @echo " - setup env $(ENV_FILE)"
    $(eval include $(ENV_FILE))
    $(eval export)
endef

## Version number of linter
LINT_VERSION ?= 1.54.2


devEnv:
	$(call setup_env,dev)
init_dev: devEnv
	docker network create -d bridge postgres --subnet 172.30.0.0/24
	docker-compose -f docker/docker-compose.yaml up -d
	sudo apt install -y postgresql-client
	PGPASSWORD=$(DB_PASSWORD) psql -h localhost -U $(DB_USER) -w -c "create database $(DB_NAME);"

create_migration: devEnv
	@read -p "Enter migration version:" version; \
	migrate create -ext sql -dir migrations -seq $$version

build:
	export CGO_ENABLED=0 && $(GO) build -mod=vendor -ldflags "-s -w"

# Утилиты
mocks: ## Generate mocks for unit tests
	go generate ./...

install-lint: ## Install linter
ifndef IS_LINT_INSTALLED
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./bin v$(LINT_VERSION)
endif
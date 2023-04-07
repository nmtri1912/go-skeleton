#!make
include .conf
export

PROJECT_NAME=server
API_CONTAINER:=$(PROJECT_NAME)-server
DB_CONTAINER_NAME=$(PROJECT_NAME)-db
DB_IMG_NAME=$(PROJECT_NAME)-db-img
PKG := "$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

# Self documented Makefile
# http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

help: ## Show list of make targets and their description
	@grep -E '^[/%.a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL:= help

run:
	go run main.go

# run: 
# 	go run cmd/server/main.go \
# 		-jwt 'test jwt' \
# 		-redis "redis://localhost:6379" \
# 		-db "postgres://postgres:postgres@localhost:54322/rmbg?sslmode=disable"


run_watch:
	CompileDaemon \
	    -exclude-dir=.git \
        -exclude-dir=.tmp \
        -command=./main \
        -build="go build cmd/server/main.go"

.PHONY: test
test: ## Generate mock and run all test. To run specified tests, use `./scripts/test.sh <pattern>`)
	@scripts/test.sh $*

.PHONY: lint
lint: ## Run linter
	@scripts/lint.sh

.PHONY: build
build/linux: ## Build server for linux
	env GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o ./bin/server-linux ./cmd/server/

build/mac: ## Show build.sh help for building binnary package under cmd
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server/

gen: ## Generate models using sqlboiler following pre-defined templates
	sqlboiler psql --wipe --add-soft-deletes --templates ./db/templates

db/migrate: ## Migrate database structure
	@scripts/migrate.sh up

db/up: ## Apply all the migration to the latest version to the local database
	@make db/migrate

db/down: ## Remove every in the database! (only for DEV)
	@scripts/migrate.sh down 

db/drop: ## Remove every in the database! (only for DEV)
	@scripts/migrate.sh drop -f

db/reset: ## Remove everything and recreate the database! (only for DEV)
	@echo y | make db/drop
	@make db/up

db/connect: ## Connect to the database
	pgcli ${DB_URL}

db/%: ## Run other migrate commands
	@scripts/migrate.sh $*


prod/build: 
	docker build -f ./docker/server.dockerfile -t ${SERVER_IMG}:${VERSION} .

docker/build: ## Build docker compose
	docker-compose build

docker/up: ## Run docker compose
	docker-compose -f ./docker/docker-compose.yaml up -d

docker/down: ## Stop docker compose
	docker-compose -f ./docker/docker-compose.yaml down

docker/clean:
	@make docker/down
	@make volume/rm

volume/rm:
	docker volume rm docker_server

redis/flush: ## Flush alll redis cache
	redis-cli -p 7100 FLUSHALL
	redis-cli -p 7101 FLUSHALL
	redis-cli -p 7102 FLUSHALL

redis/connect: ## Connect to redis master node 
	redis-cli -p 7000

setup: ## Perform setup script, install necessary plugins/tools
	@scripts/setup.sh

dep: ## Get the dependencies
	@go get -v -d ./...

clean: ## Remove previous build
	@rm -f $(PROJECT_NAME)

doc:
	scp ./docs/swagger.yml iServer:/etc/nginx/html/docs/engineer/api/server/swagger.yml

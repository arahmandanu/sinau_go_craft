SHELL         = /bin/sh

APP_NAME      = belajar-go-craft
VERSION      := $(shell git describe --always --tags)
GIT_COMMIT    = $(shell git rev-parse HEAD)
GIT_DIRTY     = $(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE    = $(shell date '+%Y-%m-%d-%H:%M:%S')
SQUAD         = "chat"
CGO_ENABLED   = 0
GOARCH		  = amd64
GOOS		  = $(shell uname -s)

.PHONY: default
default: help

.PHONY: help
help:
	@echo 'Management commands for ${APP_NAME}:'
	@echo
	@echo 'Usage:'
	@echo '    make name                  Get APP Name.'
	@echo '    make lint				  Run static linter on a compiled project.'
	@echo '    make mocks                 Generate/Update all mock files.'
	@echo '    make test                  Run tests on a compiled project.'
	@echo '    make sec			  		  Run security checks on a compiled project.'
	@echo '    make coverage			  See coverage detail.'
	@echo '    make run ARGS=             Run with supplied arguments.'
	@echo '    make build                 Compile the project.'
	@echo '    make clean                 Clean the directory tree.'
	@echo '    make prepare       		  Prepare your branch before make PR.'
	@echo '    make swag-fmt       		  Format swagger comment.'
	@echo '    make swag-gen       		  Generate swagger documentation REST API.'
	@echo '    make migrate-up       	  Run migration up.'
	@echo '    make migrate-down       	  Rollback the migration step 1'
	@echo

.PHONY: get-app-name
get-app-name:
	@echo ${APP_NAME}

.PHONY: lint
lint:
	@echo "Check linter with staticcheck"
	staticcheck ./...

.PHONY: run
run: build
	@echo "Running ${APP_NAME} ${VERSION}"
	bin/${APP_NAME} ${ARGS}

.PHONY: build
build:
	@echo "Building ${APP_NAME} ${VERSION}"
	go build -ldflags "-w -s -X github.com/arahmandanu/sinau_go_craft/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X github.com/arahmandanu/sinau_go_craft/version.Version=${VERSION} -X github.com/arahmandanu/sinau_go_craft/version.Environment=${APP_ENV} -X github.com/arahmandanu/sinau_go_craft/version.BuildDate=${BUILD_DATE}" -o bin/${APP_NAME} -trimpath .

.PHONY: prepare
prepare: lint sec test build
	@echo "Your works ready to reviewed! Go to make the PR."

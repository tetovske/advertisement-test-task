.PHONY: all

include .env
export

PROJ_PATH := ${CURDIR}
DOCKER_PATH := ${PROJ_PATH}/docker

APP=advertisement
MIGRATION_TOOL=goose
MIGRATIONS_DIR=./db/migrations

BASIC_IMAGE=dep

build:
	go build -o .bin/${APP} cmd/${APP}/main.go
	chmod ugo+x .bin/${APP}

build-docker:
	sudo rm -rf .database/
	docker build -t ${APP}-image -f ${DOCKER_PATH}/${APP}.Dockerfile .

app-setup-and-up: build-docker app-up

app-up: build
	docker-compose up

all: app-setup-and-up

app-bash:
	docker-compose run --rm --no-deps --name ${APP}-service ${APP} bash

app-up-local: build
	./.bin/advertisement

db-bash:
	docker-compose run --rm --no-deps --name ${APP}-db db ash

goose-init:
	go build -o .bin/goose cmd/${MIGRATION_TOOL}/main.go
	chmod ugo+x .bin/${MIGRATION_TOOL}

db-up:
	docker-compose run --rm --no-deps --name ${APP}-db db ash

db-migration-create: goose-init
	if [ -z ${lang} ] ; \
	then \
		goose -dir=${MIGRATIONS_DIR} create ${name} sql ; \
	else \
	  	goose -dir=${MIGRATIONS_DIR} create ${name} ${lang} ; \
	fi ;

db-migrate-status: goose-init
	docker-compose run --rm ${APP} .bin/goose -dir ${MIGRATIONS_DIR} postgres \
		"user=${POSTGRES_USER} host=${POSTGRES_HOST} port=${POSTGRES_PORT} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=${POSTGRES_SSL}" status

db-migrate-up: goose-init
	docker-compose run --rm ${APP} .bin/goose -dir ${MIGRATIONS_DIR} postgres \
        "user=${POSTGRES_USER} host=${POSTGRES_HOST} port=${POSTGRES_PORT} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=${POSTGRES_SSL}" up

db-migrate-down: goose-init
	docker-compose run --rm ${APP} .bin/goose -dir ${MIGRATIONS_DIR} postgres \
        "user=${POSTGRES_USER} host=${POSTGRES_HOST} port=${POSTGRES_PORT} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=${POSTGRES_SSL}" down

packages-tidy:
	go mod tidy
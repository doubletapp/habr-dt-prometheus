all: build down up

env:
	cp .env.example .env

build:
	docker-compose build

down:
	docker-compose down

up:
	docker-compose up -d

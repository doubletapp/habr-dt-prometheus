all: build down up

env:
	if [ ! -f .env ]; then cp .env.example .env; fi;

build:
	docker-compose build

down:
	docker-compose down

up:
	docker-compose up -d

all: up

SVCS=

.PHONY: up
up:
	docker-compose up -d ${SVCS}

.PHONY: build
build:
	docker-compose build --force-rm --no-cache ${SVCS}

.PHONY: down
down:
	docker-compose down ${SVCS}

.PHONY: logs
logs:
	docker-compose logs ${SVCS}

.PHONY: stop
stop:
	docker-compose stop ${SVCS}

.PHONY: format
format:
	docker-compose exec blog-ui npm run lint --format

.PHONY: npm-install
npm-install:
	docker-compose exec blog-ui npm install $(ARGS)

.PHONY: clean-db
clean-db:
	rm -ifrv ./backend/db/data/mysql/*

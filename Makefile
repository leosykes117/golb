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

.PHONY: clean-db
clean-db:
	yes y | rm -vrf ./backend/db/data/primary1/*
	yes y | rm -vrf ./backend/db/data/primary2/*

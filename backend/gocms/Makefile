all: bin/example
test: lint unit-test

PLATFORM=local

.PHONY: bin/example
bin/example:
	@docker buildx build . -f Dockerfile.dev --target bin \
	--output bin/ \
	--platform ${PLATFORM}

.PHONY: unit-test
unit-test:
	@docker build . --target unit-test

.PHONY: unit-test-coverage
unit-test-coverage:
	@docker build . --target unit-test-coverage \
	--output coverage/
	cat coverage/cover.out

.PHONY: lint
lint:
	@docker build . --target lint

.PHONY: debug
debug:
	@echo Construyendo imagen de ubuntu
	@docker build . \
	-t debug-app \
	-f Dockerfile.debug
.PHONY: init
# init env
init:
	go install github.com/air-verse/air@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: generate
# generate
generate:
	go generate ./...
	wire ./...
	go mod tidy

.PHONY: run
# run
run:
	go run ./...

.PHONY: air
# air
air:
	air -c .air.toml

.PHONY: air-win
# air for windows
air-win:
	air -c .air.win.toml

.PHONY: build
# build
build:
	mkdir -p bin/ & go build -o ./bin/ ./...

.PHONY: swagger
swagger:
	swag init -g cmd/server/main.go --output docs

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [targe]'
	@echo ''
	@echo 'Targets'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

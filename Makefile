.SILENT:

.DEFAULT_GOAL = help

COLOR_RESET = \033[0m
COLOR_GREEN = \033[32m
COLOR_YELLOW = \033[33m

PROJECT_NAME = `basename $(PWD)`

SHELL := /bin/bash

## prints this help
help:
	printf "${COLOR_YELLOW}\n${PROJECT_NAME}\n\n${COLOR_RESET}"
	awk '/^[a-zA-Z\-\_0-9\.%]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "${COLOR_GREEN}$$ make %s${COLOR_RESET} %s\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)
	printf "\n"

## build-image
build-image:
	docker build . -t miguelsmuller/api-thex-products

## run-image
run-image:
	docker run --name api-thex-products \
			--network local \
			-p 3002:3002 \
			-e API_ADDRESS="0.0.0.0:3002" \
			-d -it \
			miguelsmuller/api-thex-products

## stop-image
stop-image:
	docker container stop api-thex-products && \
	docker rm -f api-thex-products

## publish-image
publish-image:
	docker tag api-thex-products miguelsmuller/api-thex-products:1.0.0

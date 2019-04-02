.PHONY: all help check setup test

all: help ## same as `help`

help: ## show this help
	@sed -n -e '/@sed/d' -e 's/:.*##/:/p' Makefile

check: ## check whether the repository is in a good state
	! git status --porcelain -unormal | grep .
	cd server && go vet && go test -race

setup: client/node_modules .git/hooks/pre-push ## do first-time setup
	cd server && go install

test: ## run tests
	cd server && go test -race

.git/hooks/pre-push: ## a git hook to execute `make check`
	mkdir -p .git/hooks
	touch .git/hooks/pre-push
	chmod +x .git/hooks/pre-push
	echo '#!/bin/sh' >> .git/hooks/pre-push
	echo 'exec make check' >> .git/hooks/pre-push

client/node_modules: ## the client deps
	cd client && npm install

build: client/build server/local ## the entire project
	rm -rf build
	cp -R client/build build
	cp server/local build

build-heroku: client/build server/heroku ## the entire project, for heroku
	rm -rf build-heroku
	cp -R client/build build-heroku
	cp server/heroku build-heroku

client/build: ## the optimized client build
	cd client && npm run build

server/local: server/*.go ## the local server program
	cd server && go build -o local

server/heroku: server/*.go ## the heroku server program
	cd server && GOARCH=arm GOOS=linux go build -o heroku

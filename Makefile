.PHONY: all help check setup

all: help ## same as `help`

help: ## show this help
	@sed -n -e '/@sed/d' -e 's/:.*##/:/p' Makefile

check: ## check whether the repository is in a good state
	! git status --porcelain -unormal | grep .
	cd server && go vet && go build && go test -race

setup: ## do first-time setup
	mkdir -p .git/hooks
	rm -f .git/hooks/pre-push
	touch .git/hooks/pre-push
	chmod +x .git/hooks/pre-push
	echo '#!/bin/sh' >> .git/hooks/pre-push
	echo 'exec make check' >> .git/hooks/pre-push
	cd client && npm install
	cd server && go install

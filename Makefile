.PHONY: all help check setup

all: help ## same as `help`

help: ## show this help
	@sed -n -e '/@sed/d' -e 's/:.*##/:/p' Makefile

check: ## check whether the repository is in a good state
	! git status --porcelain -unormal | grep .
	cd server && go vet && go build && go test -race

setup: client/node_modules .git/hooks/pre-push ## do first-time setup
	cd server && go install

.git/hooks/pre-push: ## a git hook to execute `make check`
	mkdir -p $(dir $@)
	touch $@
	chmod +x $@
	echo '#!/bin/sh' >> $@
	echo 'exec make check' >> $@

client/node_modules: ## the client deps
	cd $(dir $@) && npm install

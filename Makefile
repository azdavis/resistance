.PHONY: all check setup

all: check

check:
	! git status --porcelain -unormal | grep .
	cd server && go vet && go build && go test -race

setup:
	mkdir -p .git/hooks
	rm -f .git/hooks/pre-push
	ln -s ../../scripts/check .git/hooks/pre-push
	cd client && npm install
	cd server && go install

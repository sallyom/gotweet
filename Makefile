.DEFAULT_GOAL := install

.PHONY: install
install:
	go install github.com/sallyom/gotweet/cmd/gotweet

.PHONY: local
release:
	docker build --rm -t gotweet:local .

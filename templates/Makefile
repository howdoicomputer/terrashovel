.PHONY: test

default: | build test

build:
	docker build . -t foobar

test:
	docker run --rm -v ~/.aws:/root/.aws foobar bundle exec kitchen test

env:
	docker run --rm -v ~/.aws:/root/.aws -it foobar /bin/bash

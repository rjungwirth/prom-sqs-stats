.PHONY: build run
DEFAULT_TARGET: build

VERSION := $(shell git rev-parse HEAD)
IMAGE_NAME := wmgaca/prom-sqs-stats

build:
	docker build -t $(IMAGE_NAME):$(VERSION) .
	docker tag $(IMAGE_NAME):$(VERSION) $(IMAGE_NAME):latest

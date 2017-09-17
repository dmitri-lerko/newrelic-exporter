#!/usr/bin/make -f

VERSION=$(shell git describe --tags --always)
IMAGE=previousnext/newrelic-exporter

release: build push

build:
	docker build -t ${IMAGE}:${VERSION} .

push:
	docker push ${IMAGE}:${VERSION}


.PHONY: all

all: build

build:
	go build
	yarn install

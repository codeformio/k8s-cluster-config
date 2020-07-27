.PHONY: build
build:
	docker build . -t constraint-template-builder
	docker run -v $(dir $(realpath $(firstword $(MAKEFILE_LIST)))):/work constraint-template-builder

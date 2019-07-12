# Licensed Materials - Property of IBM
# © IBM Corp. 2019

TIMESTAMP := $(shell date +%s)
VERSION := v1.03-$(TIMESTAMP)



build:
	mkdir -p out/$(VERSION)
	cd pkg/cmd ; env GOOS=linux go build -o ../../out/$(VERSION)/api-profiler-linux  -ldflags "-X main.version=$(VERSION)"
	cd pkg/cmd ; env GOOS=windows go build -o ../../out/$(VERSION)/api-profiler-windows  -ldflags "-X main.version=$(VERSION)"
	cd pkg/cmd ; env GOOS=darwin  go build -o ../../out/$(VERSION)/api-profiler-osx -ldflags "-X  main.version=$(VERSION)"

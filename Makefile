# Licensed Materials - Property of IBM
# © IBM Corp. 2019

TIMESTAMP := $(shell date +%s)
VERSION := v1.01-$(TIMESTAMP)



build:
	mkdir -p out/$(VERSION)
	cd pkg/cmd ; env GOOS=linux go build -o ../../out/$(VERSION)/api-profiler-linux-$(VERSION)
	cd pkg/cmd ; env GOOS=windows go build -o ../../out/$(VERSION)/api-profiler-windows-$(VERSION)
	cd pkg/cmd ; env GOOS=darwin  go build -o ../../out/$(VERSION)/api-profiler-osx-$(VERSION)

TIMESTAMP := $(shell date +%s)
VERSION := v1.01-$(TIMESTAMP)



build:
	mkdir ../out/$(VERSION)
	cd cmd ; env GOOS=linux go build -o ../../out/$(VERSION)/api-profiler-linux-$(VERSION)
	cd cmd ; env GOOS=windows go build -o ../../out/$(VERSION)/api-profiler-windows-$(VERSION)
	cd cmd ; env GOOS=darwin  go build -o ../../out/$(VERSION)/api-profiler-osx-$(VERSION)
	cd ..
test: format
	go run *go

TIMESTAMP := $(shell date +%s)
VERSION := v0.01-$(TIMESTAMP)

format:
	go fmt
build: format
	mkdir out/$(VERSION)
	env GOOS=linux go build -o out/$(VERSION)/api-profiler-linux-$(VERSION)
	env GOOS=windows go build -o out/$(VERSION)/api-profiler-windows-$(VERSION)
	env GOOS=darwin  go build -o out/$(VERSION)/api-profiler-osx-$(VERSION)

test: format
	go run *go

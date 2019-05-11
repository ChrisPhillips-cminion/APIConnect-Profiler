TIMESTAMP := $(shell date +%s)
VERSION := v1.01-$(TIMESTAMP)



build:
	mkdir out/$(VERSION)
	env GOOS=linux go build -o out/$(VERSION)/api-profiler-linux-$(VERSION)
	env GOOS=windows go build -o out/$(VERSION)/api-profiler-windows-$(VERSION)
	env GOOS=darwin  go build -o out/$(VERSION)/api-profiler-osx-$(VERSION)

test: format
	go run *go

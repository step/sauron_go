PKGS := $(shell go  list ./... | grep -v /vendor)

sauron:
	CGO_ENABLED=0 go build -o bin/sauron ./pkg/main/

.PHONY: sauron_stripped
sauron_stripped:
	go build -o bin/sauron -ldflags="-s -w" ./pkg/main/

.PHONY: sauron_compressed
sauron_compressed: sauron_stripped
	upx bin/sauron

test: 
	go test -cover $(PKGS)
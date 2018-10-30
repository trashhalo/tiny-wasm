current_dir=$(shell pwd)
pkg=github.com/trashhalo/tiny-wasm

deps:
	docker build . -t tinygo
	go build -o http ./serve

build:
	docker run --rm -v $(current_dir):/go/src/$(pkg) tinygo \
	build -o /go/src/$(pkg)/wasm.wasm -target wasm $(pkg)
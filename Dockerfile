FROM tinygo/tinygo
ADD wasm.syms /go/src/github.com/aykevl/tinygo/targets/wasm.syms
ADD ./tinygo/src /go/src/github.com/aykevl/tinygo/src
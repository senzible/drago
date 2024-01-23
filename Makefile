clean: 
	find ./ -name '*.wasm*' -delete 

go-wasm: 
	GOOS=js GOARCH=wasm go build -o ./examples/sandbox/web/main.wasm ./examples/sandbox/wasm/main.go
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" ./examples/sandbox/web/

tiny-wasm: 
	tinygo build -target wasm -o ./web/main.wasm ./wasm/main.go 
	cp "$(shell tinygo env TINYGOROOT)/targets/wasm_exec.js" ./web/

tiny-wasm-opt: 
	tinygo build -target wasm -o ./examples/sandbox/web/main.wasm.tiny ./examples/sandbox/wasm/main.go 
	cp "$(shell tinygo env TINYGOROOT)/targets/wasm_exec.js" ./examples/sandbox/web/
	wasm-opt -Oz -o ./examples/sandbox/web/main.wasm ./examples/sandbox/web/main.wasm.tiny 
	rm ./examples/sandbox/web/main.wasm.tiny
clean: 
	find ./ -name '*.wasm*' -delete 

go-wasm: 
	GOOS=js GOARCH=wasm go build -o ./web/main.wasm ./wasm/main.go
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" ./web/

tiny-wasm: 
	tinygo build -target wasm -o ./web/main.wasm ./wasm/main.go 
	cp "$(shell tinygo env TINYGOROOT)/targets/wasm_exec.js" ./web/

tiny-wasm-opt: 
	tinygo build -target wasm -o ./web/main.wasm.tiny ./wasm/main.go 
	cp "$(shell tinygo env TINYGOROOT)/targets/wasm_exec.js" ./web/
	wasm-opt -Oz -o ./web/main.wasm ./web/main.wasm.tiny 
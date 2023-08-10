clean: 
	find ./ -name '*.wasm*' -delete 

go-wasm: 
	GOOS=js GOARCH=wasm go build -o ./web/main.wasm ./wasm/main.go
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" ./web/
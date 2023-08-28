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
	rm ./web/main.wasm.tiny

senzible-cloud:
	tinygo build -target wasm -o ./senzible.cloud/main.wasm.tiny ./senzible.cloud.wasm/main.go 
	cp "$(shell tinygo env TINYGOROOT)/targets/wasm_exec.js" ./senzible.cloud/
	wasm-opt -Oz -o ./senzible.cloud/main.wasm ./senzible.cloud/main.wasm.tiny 
	rm ./senzible.cloud/main.wasm.tiny
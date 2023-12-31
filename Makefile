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

senzible-cloud:
	tinygo build -target wasm -o ./senzible.cloud/main.wasm.tiny ./senzible.cloud.wasm/main.go 
	cp "$(shell tinygo env TINYGOROOT)/targets/wasm_exec.js" ./senzible.cloud/
	wasm-opt -Oz -o ./senzible.cloud/main.wasm ./senzible.cloud/main.wasm.tiny 
	rm ./senzible.cloud/main.wasm.tiny

senzible-design:
	tinygo build -target wasm -o ./senzible.design/main.wasm.tiny ./senzible.design.wasm/
	cp "$(shell tinygo env TINYGOROOT)/targets/wasm_exec.js" ./senzible.design/
	wasm-opt -Oz -o ./senzible.design/main.wasm ./senzible.design/main.wasm.tiny
	rm ./senzible.design/main.wasm.tiny

senzible-tech:
	tinygo build -target wasm -o ./senzible.tech/main.wasm.tiny ./senzible.tech.wasm/main.go
	cp "$(shell tinygo env TINYGOROOT)/targets/wasm_exec.js" ./senzible.tech/
	wasm-opt -Oz -o ./senzible.tech/main.wasm ./senzible.tech/main.wasm.tiny
	rm ./senzible.tech/main.wasm.tiny

senzible: senzible-cloud senzible-design senzible-tech
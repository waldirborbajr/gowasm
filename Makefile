buildwasm:
	GOOS=js GOARCH=wasm go build -v -o ./assets/main.wasm ./cmd/wasm/main.go

build:
	go build -v -o ./bin/gowasm ./cmd/server/main.go 

genenv:
	cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./assets

runbin:
	./bin/gowasm

buildall: buildwasm  build
	
.PHONY: 

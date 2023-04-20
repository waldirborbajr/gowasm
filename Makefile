buildwasm:
	@echo "--------------------"
	@echo "Building WASM..."
	@echo "--------------------"
	GOOS=js GOARCH=wasm go build -v -o ./assets/gowasm.wasm ./cmd/wasm/main.go

build:
	@echo "--------------------"
	@echo "Building Binary..."
	@echo "--------------------"
	go build -v -o ./bin/gowasm ./cmd/server/main.go 

genenv:
	@echo "--------------------"
	@echo "Copying wasm_exec.js"
	@echo "--------------------"
	cp "$$(go env GOROOT)/misc/wasm/wasm_exec.js" ./assets

runbin: genenv buildwasm  build
	./bin/gowasm

.PHONY: 

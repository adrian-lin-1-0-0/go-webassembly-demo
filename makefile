build-wasm:
	GOOS=js GOARCH=wasm go build -o static/client.wasm client/client.go
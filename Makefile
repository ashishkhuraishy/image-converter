build:
	GOOS=js GOARCH=wasm go build -o  public/test.wasm  app/main.go
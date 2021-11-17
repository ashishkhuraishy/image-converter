build:
	GOOS=js GOARCH=wasm go build -o  public/main.wasm  app/main.go
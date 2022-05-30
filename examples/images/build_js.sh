# GOOS=js tinygo build  -o html/wasm.wam -target wasm ./win_js.go
GOOS=js GOARCH=wasm go build -o html/wasm.wasm ./main.go
go run ./serv.go

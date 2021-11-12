
default: build serve

test:
	go test ./pkg/... -race -covermode=atomic -coverprofile=coverage.out

build-bin:
	go mod tidy
	cd cmd/ui && CGO_ENABLED=0 go build -o ../../bin/mortgage-server
	cd cmd/wasm && GOOS=js GOARCH=wasm go build -o ../../ui/static/mortgage.wasm

init-ui:
	cd ui && npm i

build-ui:
	cd ui && npm run build

build: build-bin init-ui build-ui

serve:
	go run cmd/ui/main.go

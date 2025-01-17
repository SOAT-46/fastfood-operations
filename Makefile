.PHONY: build
build:
	make wire
	make swagger
	go mod tidy
	mkdir -p bin && go build -o ./bin -v ./...

.PHONY: wire
wire:
	go install github.com/google/wire/cmd/wire@latest
	cd fastfood-operations/cmd/main; wire

.PHONY: swagger
swagger:
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init -g ./fastfood-operations/cmd/main/main.go -o ./fastfood-operations/cmd/docs

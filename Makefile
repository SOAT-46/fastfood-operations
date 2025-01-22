.PHONY: goose-up
goose-up:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	goose -dir ./fastfood-operations/db/migrations postgres "host=localhost port=5432 user=postgres password=postgres dbname=fastfood-46 sslmode=disable" up

.PHONY: build
build:
	make wire
	make swagger
	go mod tidy
	mkdir -p bin && go build -o ./bin -v ./...

.PHONY: wire
wire:
	go install github.com/google/wire/cmd/wire@latest
	cd fastfood-operations/cmd; wire

.PHONY: swagger
swagger:
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init -g ./fastfood-operations/cmd/main.go -o ./fastfood-operations/cmd/docs

.PHONY: docker-up
docker-up:
	docker compose -f ./.dev/compose.yaml up --build
	make goose-up

.PHONY: lint
lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	cd fastfood-operations && golangci-lint run \
		--config ".golangci.yaml" \
		--color "always" \
		--timeout "10m" \
		--print-resources-usage \
		--allow-parallel-runners \
		--max-issues-per-linter 0 \
		--max-same-issues 0 ./...

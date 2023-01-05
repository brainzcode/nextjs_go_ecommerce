build:
	@go build -o bin/nextjs_go_ecommerce

run: build
	@./bin/nextjs_go_ecommerce

test:
	@go test -v ./...
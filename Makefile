# Makefile
run:
	@templ generate --watch --proxy="http://localhost:8080" --cmd="go run ./cmd/main.go"
build:
	@go build -o ./tmp/main ./cmd/main.go


sqlc:
	sqlc generate

templates:
	templ generate

run: sqlc templates
	go run ./cmd/live/main.go
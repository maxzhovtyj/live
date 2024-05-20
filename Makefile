sqlc:
	sqlc generate

templates:
	templ generate

run: sqlc templates
	go run ./cmd/live/main.go

LIVE_USER=root
LIVE_HOST=194.164.59.123
LIVE_PATH=/var/www/live

live-linux:
	GOARCH=amd64 GOOS=linux go build -o bin/live-linux-amd64 ./cmd/live/

deploy-live: live-linux
	rsync -r bin/live-linux-amd64 cmd.sh ./static $(LIVE_USER)@$(LIVE_HOST):$(LIVE_PATH)
	ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null $(LIVE_USER)@$(LIVE_HOST) "cd $(LIVE_PATH) && bash ./cmd.sh"
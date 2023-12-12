deps/up: 
	docker compose up -d

deps/down: 
	docker compose down

api/dev:
	go run main.go api

consumer/dev:
	go run main.go consumer

install:
	go mod vendor
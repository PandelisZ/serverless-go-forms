build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/json functions/json.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/urlencoded functions/urlencoded.go
deploy:
	make build
	serverless deploy
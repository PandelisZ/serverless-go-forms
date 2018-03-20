build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/function function.go
deploy:
	make build
	serverless deploy

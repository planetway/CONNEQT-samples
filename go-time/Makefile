build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/go-time main.go
	docker build -t planetway/conneqt-samples-go-time .

push:
	docker push planetway/conneqt-samples-go-time:latest

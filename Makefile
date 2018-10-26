pack:
	GOOS=linux GOARCH=amd64 go build -o dist/proxy proxy.go
	docker build -t urlencode-proxy:latest .

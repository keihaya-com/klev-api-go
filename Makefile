all:
	go build -v ./...
	go test -v -cover ./...

update-libs:
	go get -u github.com/klev-dev/kleverr@main
	go mod tidy

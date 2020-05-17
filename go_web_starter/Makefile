fixfmt:
	@echo "Fixing format"
	@gofmt -w -l .

test:
	@echo "Running Tests"
	@go test ./...

mocks:
	@mockgen -destination=$(pwd)/usecase/mocks/mock_storage.go -source=$(pwd)//usecase/storage/storage.go

compile-linux:
	@echo "Building Binaries"
	@GOOS=darwin GOARCH=amd64 go build -o=build/godownload-darwin app/main.go
	@GOOS=linux GOARCH=amd64 go build -o=build/godownload-linux app/main.go
	@GOOS=windows GOARCH=amd64 go build -o=build/godownload-windows app/main.go

clean:
	@echo "Removing existing builds"
	@-rm -rf build
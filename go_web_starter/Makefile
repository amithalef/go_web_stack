fixfmt:
	@echo "Fixing format"
	@gofmt -w -l .

test:
	@echo "Running Tests"
	@go test ./...

mocks:
    @mockgen -destination=$(pwd)/usecase/mocks/mock_storage.go -source=$(pwd)//usecase/storage/storage.go
fixfmt:
	@echo "Fixing format"
	@gofmt -w -l .

compile:
	@echo "Building Binaries"
	@GOOS=darwin GOARCH=amd64 go build -o=build/godownload-darwin app/main.go
	@GOOS=linux GOARCH=amd64 go build -o=build/godownload-linux app/main.go
	@GOOS=windows GOARCH=amd64 go build -o=build/godownload-windows app/main.go

clean:
	@echo "Removing existing builds"
	@-rm -rf build

docker:
	@docker build -t load-test .

compose:
	@echo "######### removing existing deployment ########"
	@docker-compose down
	@echo "######### building container ##################"
	@docker-compose build
	@echo "######### starting containers #################"
	@docker-compose up
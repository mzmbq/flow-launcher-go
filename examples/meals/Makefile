.PHONY: build
build:
	@GOOS=windows GOARCH=amd64 go build -o ./bin/meals/meals.exe -v ./meals.go
	@cp ./assets/* ./bin/meals

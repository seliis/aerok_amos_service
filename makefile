webapp:
	@rm -rf public && rm -rf app/public
	@cd app && flutter clean && flutter build web --output=public
	@mv app/public .

server:
	@go build -o main.exe src/main.go

build:
	@make webapp
	@make server

run:
	@go run src/main.go
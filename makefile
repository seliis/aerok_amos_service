run:
	@go run main.go

test:
	@go test -v ./...

db:
	@go run github.com/steebchen/prisma-client-go db push --schema schema.prisma
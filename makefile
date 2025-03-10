run:
	@go run main.go

test:
	@go test -v ./...

db:
	@go run github.com/steebchen/prisma-client-go db push --schema schema.prisma

release:
	@make build-clear
	@make build-server
	@make build-app
	@cp -r public build/public
	@cp prisma/database.db build/prisma/database.db
	@cp cert.pem build/cert.pem
	@cp key.pem build/key.pem

build-clear:
	@rm -rf build public app/build
	@mkdir -p build/prisma public

build-server:
	@go build -o build/main.exe main.go
	@cp settings.toml build/settings.toml

build-app:
	@cd app && flutter build web --no-tree-shake-icons --release
	@cp -r app/build/web/* public
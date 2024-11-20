COPY_FILES = cert.pem cert.key config.toml
TARGET_DIR = public
BUILD_DIR = build
BINARY = main.exe

webapp:
	@rm -rf $(BUILD_DIR) && rm -rf app/$(TARGET_DIR) && mkdir $(BUILD_DIR)
	@cd app && flutter clean && flutter build web --output=$(TARGET_DIR)
	@mv app/$(TARGET_DIR) $(BUILD_DIR)

server:
	@go build -o $(BUILD_DIR)/$(BINARY) src/main.go

package:
	@make webapp
	@make server
	@cp $(COPY_FILES) $(BUILD_DIR)

run:
	@go run src/main.go
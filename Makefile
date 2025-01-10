BUILD_DIR := build
APP_NAME := wallet
OS_ARCHS := \
    linux/amd64 \
    linux/arm64 \
    darwin/amd64 \
    darwin/arm64 \

GOMOD := go mod tidy && go mod vendor
GO_BUILD := go build

.PHONY: all
all: build-prod

.PHONY: gomod
gomod:
	$(GOMOD)

.PHONY: build-prod
build-prod: clean gomod
	@echo "Building for multiple platforms..."
	@mkdir -p $(BUILD_DIR)
	@for os_arch in $(OS_ARCHS); do \
		os=$$(echo $$os_arch | cut -d'/' -f1); \
		arch=$$(echo $$os_arch | cut -d'/' -f2); \
		output=$(BUILD_DIR)/$(APP_NAME)-$$os-$$arch; \
		echo "Building $$os/$$arch..."; \
		GOOS=$$os GOARCH=$$arch $(GO_BUILD) -o $$output ./cmd; \
	done

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)


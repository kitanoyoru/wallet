SOLC = solc
ABIGEN = abigen

CONTRACTS_DIR = solidity
BUILD_DIR = build
GEN_DIR = gen

GO_PKG = contracts 

CONTRACTS      = $(CONTRACTS_DIR)/Wallet.sol $(CONTRACTS_DIR)/Allowance.sol
COMPILED_FILES = $(BUILD_DIR)/Wallet.json $(BUILD_DIR)/Allowance.json
GO_FILES       = $(GEN_DIR)/Wallet.go $(GEN_DIR)/Allowance.go

all: $(COMPILED_FILES) $(GO_FILES)

$(BUILD_DIR)/%.json: $(CONTRACTS_DIR)/%.sol
	@mkdir -p $(BUILD_DIR)
	$(SOLC) --optimize --abi --bin --overwrite -o $(BUILD_DIR) $<
	@echo "Compiled $< to $(BUILD_DIR)/$*.json"

$(GEN_DIR)/%.go: $(BUILD_DIR)/%.json
	@mkdir -p $(GEN_DIR)
	$(ABIGEN) \
		--bin=$(BUILD_DIR)/$*.bin \
		--abi=$(BUILD_DIR)/$*.abi \
		--pkg=$(GO_PKG) \
		--out=$(GEN_DIR)/$*.go
	@echo "Generated Go interface for $* in $(GEN_DIR)/$*.go"

clean:
	rm -rf $(BUILD_DIR) $(GEN_DIR)

.PHONY: all clean

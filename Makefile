VERSION := 0.0.1
COMMIT := $(shell git log -1 --format='%H')

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=Human \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=humans \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)'
DOCKER := $(shell which docker)
DOCKER_BUF := $(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace bufbuild/buf
PROJECT_NAME = $(shell git remote get-url origin | xargs basename -s .git)

all: install

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/humansd-manager
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/humansd

build: go.sum clean
	go build -mod=mod $(BUILD_FLAGS) -o build/humansd-manager ./cmd/humansd-manager
	go build -mod=mod $(BUILD_FLAGS) -o build/humansd ./cmd/humansd

build-linux:
	GOOS=linux GOARCH=amd64 $(MAKE) build

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	GO111MODULE=on go mod verify

# look into .golangci.yml for enabling / disabling linters
lint:
	@echo "--> Running linter"
	@golangci-lint run
	@go mod verify

# devnet

devnet: clean install devnet-prepare devnet-start

devnet-prepare:
	./scripts/prepare-devnet.sh

devnet-start:
	DAEMON_NAME=humansd DAEMON_HOME=~/.humans DAEMON_ALLOW_DOWNLOAD_BINARIES=true DAEMON_RESTART_AFTER_UPGRADE=true \
    humansd start --pruning="nothing" --inv-check-period 5

# Clean up the build directory
clean:
	rm -rf build/


# Localnet

# Build nodes using Docker
build-docker:
	$(MAKE) -C networks/local

# Run a 4-node testnet locally
localnet-start: build-linux localnet-stop
	@if ! [ -f build/node0/humansd/config/genesis.json ]; then docker run --rm -v $(CURDIR)/build:/humansd:Z lottery/core testnet --v 4 -o . --starting-ip-address 192.168.10.2 --keyring-backend=test --chain-id test ; fi
	./scripts/import-localnet-seeds.sh
	docker-compose up

# Stop testnet
localnet-stop:
	docker-compose down

localnet: clean build-linux build-docker localnet-start

###############################################################################
###                                Protobuf                                 ###
###############################################################################

protoVer=v0.3
protoImageName=tendermintdev/sdk-proto-gen:$(protoVer)
containerProtoGen=$(PROJECT_NAME)-proto-gen-$(protoVer)
containerProtoGenAny=$(PROJECT_NAME)-proto-gen-any-$(protoVer)
containerProtoGenSwagger=$(PROJECT_NAME)-proto-gen-swagger-$(protoVer)
containerProtoFmt=$(PROJECT_NAME)-proto-fmt-$(protoVer)

proto-all: proto-format proto-lint proto-gen

proto-gen:
	docker run --rm -v $(CURDIR):/workspace --workdir /workspace bharvest/liquidity-proto-gen sh ./scripts/protocgen.sh
	go mod tidy

# This generates the SDK's custom wrapper for google.protobuf.Any. It should only be run manually when needed
proto-gen-js:
	@echo "Generating Protobuf Typescript"
	bash ./scripts/protocgen-js.sh

proto-swagger-gen:
	@echo "Generating Protobuf Swagger"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGenSwagger}$$"; then docker start -a $(containerProtoGenSwagger); else docker run --name $(containerProtoGenSwagger) -v $(CURDIR):/workspace --workdir /workspace $(protoImageName) \
		sh ./scripts/protoc-swagger-gen.sh; fi

proto-format:
	@echo "Formatting Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoFmt}$$"; then docker start -a $(containerProtoFmt); else docker run --name $(containerProtoFmt) -v $(CURDIR):/workspace --workdir /workspace tendermintdev/docker-build-proto \
		find ./ -not -path "./third_party/*" -name "*.proto" -exec clang-format -i {} \; ; fi


proto-lint:
	@$(DOCKER_BUF) lint --error-format=json

proto-check-breaking:
	@$(DOCKER_BUF) breaking --against $(HTTPS_GIT)#branch=master

# Create log files
log-files:
	sudo mkdir -p /var/log/humansd && sudo touch /var/log/humansd/humansd.log && sudo touch /var/log/humansd/humansd_error.log

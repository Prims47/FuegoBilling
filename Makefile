.PHONY: vendor

vendor:
	go mod vendor && go mod tidy

.PHONY: build_macos

build_macos:
	GO111MODULE=on GOOS=darwin GOARCH=amd64 go build -o bin/fuegoBilling ./cmd/fuego_billing


.PHONY: docker_mockgen_build

docker_mockgen_build:
	docker build -t fuegobilling/mockgen .

.PHONY: gen_mock

gen_mock: 
	docker run -v $(PWD):/app -w /app fuegobilling/mockgen mockgen -source=internal/adapter/account_adapter_interface.go -destination=internal/repository/mock/account_adapter_mock.go -package=repository
	docker run -v $(PWD):/app -w /app fuegobilling/mockgen mockgen -source=internal/adapter/customer_adapter_interface.go -destination=internal/repository/mock/customer_adapter_mock.go -package=repository

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
	docker run -v $(PWD):/app -w /app fuegobilling/mockgen mockgen -source=internal/adapter/service_adapter_interface.go -destination=internal/repository/mock/service_adapter_mock.go -package=repository

	docker run -v $(PWD):/app -w /app fuegobilling/mockgen mockgen -source=internal/repository/account_repository_interface.go -destination=cmd/fuego_billing/mock/account_repository_mock.go -package=cmd
	docker run -v $(PWD):/app -w /app fuegobilling/mockgen mockgen -source=internal/repository/customer_repository_interface.go -destination=cmd/fuego_billing/mock/customer_repository_mock.go -package=cmd
	docker run -v $(PWD):/app -w /app fuegobilling/mockgen mockgen -source=internal/repository/service_repository_interface.go -destination=cmd/fuego_billing/mock/service_repository_mock.go -package=cmd
	docker run -v $(PWD):/app -w /app fuegobilling/mockgen mockgen -source=internal/services/format_float.go -destination=cmd/fuego_billing/mock/format_float_mock.go -package=cmd
	docker run -v $(PWD):/app -w /app fuegobilling/mockgen mockgen -source=internal/services/format_int.go -destination=cmd/fuego_billing/mock/format_int_mock.go -package=cmd

.PHONY: unit_test

unit_test:
	go test -coverprofile=coverage.out -v ./... 

.PHONY: test_coverage_html

test_coverage_html:
	go tool cover -html=coverage.out
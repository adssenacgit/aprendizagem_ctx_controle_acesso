.PHONY: all check-dependencies

.DEFAULT_GOAL := help


check-dependencies: ## Check dependencies
	@docker --version
	@openssl version



help: ## Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?# "}; {printf "\033[36m%-40s\033[0m %s\n", $$1, $$2}'

generate_keys: ## Generate private and public key
	@openssl genrsa -out ./certs/private.pem 3072
	@openssl rsa -in ./certs/private.pem -pubout -out ./certs/public.pem
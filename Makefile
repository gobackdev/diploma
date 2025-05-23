MAIN_FILE := cmd/gophermart/main.go
ENV_FILE := cmd/gophermart/.env

dev:
	@if [ -f $(ENV_FILE) ]; then \
		export $$(grep -v '^#' $(ENV_FILE) | xargs); \
	fi && go run $(MAIN_FILE) -a :8083
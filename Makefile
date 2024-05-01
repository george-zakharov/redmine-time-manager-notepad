BINARY_NAME=notepad

all: build start ## Build server and run it

run: ## Simple Go run with PROD mode
	@echo "  >  Running server in PROD mode..."
	go run main.go

runDev: ## Simple Go run with DEV mode
	@echo "  >  Running server in DEV mode..."
	go run main.go -mode=1

build: ## Build Go Server to run
	@echo "  >  Building server..."
	go build -o ${BINARY_NAME} main.go

dockerBuild: ## Build Go Server for Docker
	@echo "  >  Building server..."
	go build -ldflags "-s -w" -o ${BINARY_NAME} main.go

start: ## Start compiled server
	@echo "  >  Starting server..."
	./${BINARY_NAME}

clean: ## Clean development env
	@echo "  >  Cleaning Go env..."
	go clean
	rm ${BINARY_NAME}

doc: ## Generate docs
	@echo " > Generating docs..."
	swag init -g main.go

help: Makefile
	@echo " Choose a command to run in "${BINARY_NAME}":"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

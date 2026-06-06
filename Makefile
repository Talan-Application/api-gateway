SHELL         = powershell.exe
.SHELLFLAGS   = -NoProfile -NonInteractive -Command

APP_NAME  = api-gateway
BUILD_DIR = ./bin
BINARY    = $(BUILD_DIR)/$(APP_NAME).exe
MAIN      = ./cmd/

.PHONY: all build run test test-verbose fmt tidy lint clean docker-build docker-run docker-stop proto-update help

all: build

build:
	New-Item -ItemType Directory -Force -Path $(BUILD_DIR) | Out-Null; go build -ldflags="-s -w" -o $(BINARY) $(MAIN)

run:
	go run $(MAIN)

test:
	go test ./...

test-verbose:
	go test -v -race ./...

fmt:
	gofmt -s -w .

tidy:
	go mod tidy

proto-update:
	$$env:GOPROXY = "direct"; go get github.com/Talan-Application/proto-generation@master; go mod tidy

lint:
	golangci-lint run ./...

clean:
	Remove-Item -Recurse -Force -ErrorAction SilentlyContinue $(BUILD_DIR)

docker-build:
	docker build -t $(APP_NAME):latest .

docker-run:
	docker run --rm --env-file .env -p 8080:8080 --name $(APP_NAME) $(APP_NAME):latest

docker-stop:
	docker stop $(APP_NAME)

help:
	Write-Host "Usage: make <target>"; \
	Write-Host ""; \
	Write-Host "  build          Build the binary to $(BINARY)"; \
	Write-Host "  run            Run the service locally (go run)"; \
	Write-Host "  test           Run tests"; \
	Write-Host "  test-verbose   Run tests with -v -race"; \
	Write-Host "  fmt            Format source files"; \
	Write-Host "  tidy           Run go mod tidy"; \
	Write-Host "  proto-update   Fetch latest proto-generation module (GOPROXY=direct) and tidy"; \
	Write-Host "  lint           Run golangci-lint"; \
	Write-Host "  clean          Remove build artifacts"; \
	Write-Host "  docker-build   Build Docker image"; \
	Write-Host "  docker-run     Run Docker container with .env"; \
	Write-Host "  docker-stop    Stop running container"

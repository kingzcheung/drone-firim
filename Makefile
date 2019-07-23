GOCMD=go
PROJECT_NAME=firim

.PHONY: build
build:
	PROJECT_NAMECGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOPROXY=https://goproxy.io $(GOCMD) build -o ./$(PROJECT_NAME)/$(PROJECT_NAME) ./$(PROJECT_NAME)/main.go

.PHONY: mac
mac:
	GOPROXY=https://goproxy.io $(GOCMD) build -o ./$(PROJECT_NAME)/$(PROJECT_NAME) ./$(PROJECT_NAME)/main.go
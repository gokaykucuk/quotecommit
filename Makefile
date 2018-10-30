# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=bin/gitqc
INSTALL_PATH=/usr/local/

all: test build
build:
		$(GOBUILD) -o $(BINARY_NAME) -v
test:
		$(GOTEST) -v ./...
clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
run:
		$(GOBUILD) -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)

install:
		cp $(BINARY_NAME) $(INSTALL_PATH)$(BINARY_NAME)
uninstall:
		rm $(INSTALL_PATH)$(BINARY_NAME)
update_data:
		$(GOCMD) run main.go update_data
APP_NAME=bin/demo
APP_SRC=main.go

GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean


.PHONY: all
all: build

.PHONY: build
build:
	$(GOBUILD) -o $(APP_NAME) -v $(APP_SRC)

.PHONY: test
test:
	$(GOTEST)

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(APP_NAME)

.PHONY: run
run:
	$(GOBUILD) -o $(APP_NAME) -v $(APP_SRC)
	./$(APP_NAME)


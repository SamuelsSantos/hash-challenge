.PHONY: all
all: build
FORCE: ;

BIN_DIR = $(PWD)/bin

.PHONY: build

clean:
	./mvnw clean

dependencies:
	./mvnw install -U

compile:
	./mvnw compile

build: dependencies clean compile build-api

build-api: 
	./mvnw package -Dquarkus.package.type=fast-jar


docker-build: build
	@docker image build -t web-api . -f Dockerfile.fast-jar

.PHONY: all
all: build
FORCE: ;

BIN_DIR = $(PWD)/bin

.PHONY: build

clean:
	rm -rf bin/*

dependencies:
	go mod download

build: dependencies clean build-api linux-binaries

build-api: 
	go build -o ./bin/grpc-products api/grpc/main.go

linux-binaries:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o ./bin/grpc-products-linux api/grpc/main.go

fmt: ## gofmt and goimports all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

docker-build: build
	@docker image build -t grpc-products . 


server:
	go run api/grpc/main.go -port 8486

test:
	mkdir -p ./coverage
	@for d in $$(go list ./... | grep -v /domain/pb | grep -v /cmd); do go test $${d} -v -coverprofile=./coverage/tests.out; done
	go tool cover -html=./coverage/tests.out -o ./coverage/coverage-report.html
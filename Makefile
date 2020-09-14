gen:
	mkdir -p $(SERVICE)/domain/pb
	protoc --proto_path=protorepo protorepo/$(SERVICE).proto --go_out=plugins=grpc:$(SERVICE)/domain/pb

gen-discount-service:
	mkdir -p discountcalculator/domain/pb
	protoc --proto_path=protorepo protorepo/*.proto --go_out=plugins=grpc:discountcalculator/domain/pb

msgtoProduct:
	cp discount-calculator/proto/* product-list/src/main/proto

msgtoDiscountCalculator:
	cp product-list/src/main/proto/* discount-calculator/proto


docker-rebuild: build-all
	@docker-compose down 
	@docker system prune -f
	@docker volume prune -f
	@docker rmi user-service:latest product-service:latest discount-calculator-service:latest
	@docker-compose up -d --build
	@docker-compose logs -f

docker-up:
	@docker-compose up -d


docker-upbuild:
	@docker-compose up -d --build
	@docker-compose logs -f

build-all:
	$(MAKE) -C products build
	$(MAKE) -C users build
	$(MAKE) -C discountcalculator build
	$(MAKE) -C product-list build

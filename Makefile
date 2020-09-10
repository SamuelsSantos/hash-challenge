gen:
	mkdir -p $(SERVICE)/domain/pb
	protoc --proto_path=protorepo protorepo/$(SERVICE).proto --go_out=plugins=grpc:$(SERVICE)/domain/pb

gen-discount-service:
	mkdir -p discountcalculator/domain/pb
	protoc --proto_path=protorepo protorepo/*.proto --go_out=plugins=grpc:discountcalculator/domain/pb

clean:
	rm $(SERVICE)/pb/*.go

msgtoProduct:
	cp discount-calculator/proto/* product-list/src/main/proto

msgtoDiscountCalculator:
	cp product-list/src/main/proto/* discount-calculator/proto

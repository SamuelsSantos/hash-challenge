gen:
	mkdir -p $(SERVICE)/pb
	protoc --proto_path=protorepo protorepo/$(SERVICE)/*.proto --go_out=plugins=grpc:$(SERVICE)/pb

clean:
	rm $(SERVICE)/pb/*.go

msgtoProduct:
	cp discount-calculator/proto/* product-list/src/main/proto

msgtoDiscountCalculator:
	cp product-list/src/main/proto/* discount-calculator/proto

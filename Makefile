build:
	protoc -I. --go_out=plugins=micro:/Users/tattoor/source/consignment/consignment-service \
					proto/consignment/consignment.proto
	docker build -t consignment-cli .

run:
	docker run -e MICRO_REGISTRY=mdns consignment-cli
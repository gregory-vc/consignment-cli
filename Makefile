build:
	go get github.com/gregory-vc/consignment-service
	go mod vendor

run:
	docker run -e MICRO_REGISTRY=mdns consignment-cli
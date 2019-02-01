build:
		go get github.com/gregory-vc/consignment-service
		go mod vendor
		git add --all
		git diff-index --quiet HEAD || git commit -a -m 'fix'
		git push origin master
run:
		docker run -e MICRO_REGISTRY=mdns consignment-cli
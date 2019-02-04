package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/gregory-vc/consignment-service/proto/consignment"

	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
	"golang.org/x/net/context"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {

	cmd.Init()

	// Create new greeter client
	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)

	// Contact the server and print out its response.
	var token string
	log.Println(os.Args)

	token = os.Args[1]

	consignment, err := parseFile(defaultFilename)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	// Create a new context which contains our given token.
	// This same context will be passed into both the calls we make
	// to our consignment-service.
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})

	// First call using our tokenised context
	r, err := client.CreateConsignment(ctx, consignment)
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	// Second call
	getAll, err := client.GetConsignments(ctx, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}

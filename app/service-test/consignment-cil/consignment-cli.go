// consignment-cli/cli.go
package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
	pb "github.com/mongmx/mymicro/app/consignment-service/proto/consignment"
	"golang.org/x/net/context"
)

const (
	address         = "localhost:50051"
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
	file := defaultFilename
	// if len(os.Args) > 1 {
	// 	file = os.Args[1]
	// }
	var token string
	log.Println(os.Args)

	if len(os.Args) < 3 {
		log.Fatal(errors.New("Not enough arguments, expecing file and token"))
	}

	file = os.Args[1]
	token = os.Args[2]

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	// Create a new context which contains our given token.
	// This same context will be passed into both the calls we make
	// to our consignment-service.
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})

	r, err := client.CreateConsignment(ctx, consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(ctx, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}

}

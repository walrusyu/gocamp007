package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	pb "github.com/walrusyu/gocamp007/demo/api/bff/v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:5001", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewBffServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetUser(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Get User: %v", r)

	//var messageType2 *descriptorpb.DescriptorProto
	var messageType *pb.User
	fm, err := fieldmaskpb.New(messageType, "age")
	if err == nil {
		r, err = c.UpdateUser(ctx, &pb.UpdateUserRequest{
			User: &pb.User{
				Id:   &wrappers.Int32Value{Value: 1},
				Name: "ywf",
				Age:  11,
				Address: &pb.User_Address{
					Province: "sh",
					City:     "hk",
					Street:   "lc",
				},
			},
			Mask: fm,
		})
	} else {
		fmt.Printf("error:%v", err.Error())
	}

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Update User: %v", r)
}

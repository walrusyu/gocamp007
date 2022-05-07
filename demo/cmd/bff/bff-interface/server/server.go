package server

import (
	context "context"
	"fmt"
	pb "github.com/walrusyu/gocamp007/demo/api/bff/v1"
	orderpb "github.com/walrusyu/gocamp007/demo/api/order/v1"
	userpb "github.com/walrusyu/gocamp007/demo/api/user/v1"
	"github.com/walrusyu/gocamp007/demo/cmd/bff/bff-interface/client"
	"github.com/walrusyu/gocamp007/demo/config"
	cErros "github.com/walrusyu/gocamp007/demo/errors"
	"github.com/walrusyu/gocamp007/demo/internal/middleware/tracing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

var _ pb.BffServiceServer = &server{}

type server struct {
	pb.UnimplementedBffServiceServer
	config           config.BffServerConfig
	orderServiceAddr string
	userServiceAddr  string
}

func NewServer(config config.BffServerConfig) pb.BffServiceServer {
	return &server{
		config:           config,
		orderServiceAddr: fmt.Sprintf("%s:%d", config.OrderServiceIP, config.OrderServicePort),
		userServiceAddr:  fmt.Sprintf("%s:%d", config.UserServiceIP, config.UserServicePort),
	}
}

func (s *server) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.Order, error) {
	c := make(chan *pb.Order, 1)
	defer close(c)
	go func() {
		// Set up a connection to the server.
		conn, err := client.Dial(ctx,
			client.WithEndpoint("127.0.0.1:9000"),
			client.WithMiddleware(
				tracing.TracingMiddlewareOnClient(),
			),
			client.WithTimeout(2*time.Second))

		// for tracing remote ip recording
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		client := orderpb.NewOrderServiceClient(conn)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := client.Get(ctx, &orderpb.GetRequest{Id: req.GetId()})
		if err != nil {
			log.Fatalf("could not get order: %v", err)
		}
		log.Printf("Order: %v", r)
		order := &pb.Order{
			Id:          r.Id,
			Description: r.Description,
			OrderItems:  []*pb.Order_OrderItem{},
		}
		for i := range r.OrderItems {
			order.OrderItems = append(order.OrderItems, &pb.Order_OrderItem{
				Id:       r.OrderItems[i].Id,
				Offer:    r.OrderItems[i].Offer,
				Quantity: r.OrderItems[i].Quantity,
			})
		}
		c <- order
	}()
	select {
	case <-ctx.Done():
		return nil, cErros.New(500, "timeout", "can not get order")
	default:
		return <-c, nil
	}
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	c := make(chan *pb.User, 1)
	defer close(c)
	go func() {
		// Set up a connection to the server.
		conn, err := grpc.Dial(s.userServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		client := userpb.NewUserServiceClient(conn)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := client.Get(ctx, &userpb.GetRequest{Id: req.GetId()})
		if err != nil {
			log.Fatalf("could not get user: %v", err)
		}
		log.Printf("User: %v", r)

		user := &pb.User{
			Id:   r.Id,
			Name: r.Name,
			Age:  r.Age,
			Address: &pb.User_Address{
				Province: r.Address.Province,
				City:     r.Address.City,
				Street:   r.Address.Street,
			},
		}
		c <- user
	}()
	select {
	case <-ctx.Done():
		return nil, cErros.New(500, "timeout", "can not get user")
	default:
		return <-c, nil
	}
}

func (s *server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.User, error) {
	c := make(chan *pb.User, 1)
	defer close(c)
	go func() {
		// Set up a connection to the server.
		conn, err := grpc.Dial(s.userServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		client := userpb.NewUserServiceClient(conn)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := client.Update(ctx, &userpb.UpdateRequest{User: &userpb.User{
			Id:   req.User.Id,
			Name: req.User.Name,
			Age:  req.User.Age,
			Address: &userpb.User_Address{
				Province: req.User.Address.Province,
				City:     req.User.Address.City,
				Street:   req.User.Address.Street,
			},
		}, Mask: req.Mask})
		if err != nil {
			log.Fatalf("could not update user: %v", err)
		}
		log.Printf("User: %v", r)
		user := &pb.User{
			Id:   r.Id,
			Name: r.Name,
			Age:  r.Age,
			Address: &pb.User_Address{
				Province: r.Address.Province,
				City:     r.Address.City,
				Street:   r.Address.Street,
			},
		}
		c <- user
	}()
	select {
	case <-ctx.Done():
		return nil, cErros.New(500, "timeout", "can not get user")
	default:
		return <-c, nil
	}
}

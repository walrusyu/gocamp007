package server

import (
	context "context"
	"fmt"
	pb "github.com/walrusyu/gocamp007/demo/api/bff/v1"
	orderpb "github.com/walrusyu/gocamp007/demo/api/order/v1"
	userpb "github.com/walrusyu/gocamp007/demo/api/user/v1"
	"github.com/walrusyu/gocamp007/demo/config"
	cErros "github.com/walrusyu/gocamp007/demo/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"time"
)

type Server struct {
	pb.UnimplementedBffServiceServer
	config           config.BffServerConfig
	orderServiceAddr string
	userServiceAddr  string
}

func NewServer(config config.BffServerConfig) *Server {
	return &Server{
		config:           config,
		orderServiceAddr: fmt.Sprintf("%s:%d", config.OrderServiceIP, config.OrderServicePort),
		userServiceAddr:  fmt.Sprintf("%s:%d", config.UserServiceIP, config.UserServicePort),
	}
}

func (s *Server) GetOrder(ctx context.Context, req *emptypb.Empty) (*pb.Order, error) {
	c := make(chan *pb.Order, 1)
	defer close(c)
	go func() {
		// Set up a connection to the server.
		conn, err := grpc.Dial(s.orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		client := orderpb.NewOrderServiceClient(conn)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := client.Get(ctx, &emptypb.Empty{})
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

func (s *Server) GetUser(ctx context.Context, req *emptypb.Empty) (*pb.User, error) {
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
		r, err := client.Get(ctx, &emptypb.Empty{})
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

func (s *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.User, error) {
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

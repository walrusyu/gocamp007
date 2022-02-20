package service

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/walrusyu/gocamp007/week4/api/user"
	"github.com/walrusyu/gocamp007/week4/internal/biz"
)

type Server struct {
	pb.UnimplementedUserServer
	biz *biz.UserBiz
}

func (s *Server) SaveAddressBook(ctx context.Context, in *pb.SaveAddressBookRequest) (*pb.SaveAddressBookReply, error) {
	people := in.GetPeople()
	if people == nil {
		return nil, errors.New("incorrect request")
	}
	var person []*biz.Person
	for _, p := range people {
		newPerson := &biz.Person{Id: p.GetId(), Name: p.GetName(), Email: p.GetEmail()}
		for _, phoneNumber := range p.Phones {
			newPerson.PhoneNumbers = append(newPerson.PhoneNumbers, &biz.PhoneNumber{Number: phoneNumber.Number, Type: int32(phoneNumber.Type)})
		}
		person = append(person, newPerson)
	}
	err := s.biz.SaveAddressBook(person)
	if err == nil {
		return nil, errors.New("save failed")
	}
	return &pb.SaveAddressBookReply{Message: fmt.Sprintf("saved %d people", len(person))}, nil
}

func NewServer(biz *biz.UserBiz) *Server {
	return &Server{biz: biz}
}

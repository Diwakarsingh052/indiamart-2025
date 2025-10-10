package main

import (
	"context"
	"fmt"
	pb "server/gen/proto"
	"server/models"

	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func (us *UserService) Signup(ctx context.Context,
	req *pb.SignupRequest) (*pb.SignupResponse, error) {
	var u models.User

	user := req.GetUser()
	u.Name = user.GetName()
	u.Email = user.GetEmail()
	u.Password = user.GetPassword()
	u.Roles = user.GetRoles()

	validate := validator.New()

	err := validate.Struct(u)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user")
	}

	// call the business logic after validating the request
	//user.CreateUser(u)
	fmt.Println("received request and processed", u)

	return &pb.SignupResponse{Result: "User created Successfully"}, nil

}

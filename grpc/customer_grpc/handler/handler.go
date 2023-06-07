package handler

import (
	"context"
	"jet/grpc/customer_grpc/repository"
	"jet/pb"
	"log"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type CustomerServiceImp interface {
	CreateCustomer(ctx context.Context, model *pb.CreateCustomerRequest) (*pb.Customer, error)
}

type CustomerService struct {
	customerRepository repository.CustomerRepository
	pb.UnimplementedCustomerServiceServer
}

func NewCustomerrHandler(customerRepository repository.CustomerRepository) (*CustomerService, error) {
	return &CustomerService{
		customerRepository: customerRepository,
	}, nil
}

func (s *CustomerService) CreateCustomer(ctx context.Context, model *pb.CreateCustomerRequest) (*pb.Customer, error) {
	res, err := s.customerRepository.CreateCustomer(ctx, model)
	if err != nil {
		log.Println("Created failed:", err)
		return nil, err
	}

	return &pb.Customer{
		Id:        res.ID.String(),
		Name:      res.Fullname,
		Phone:     res.Phone,
		Address:   res.Address,
		Gender:    pb.Gender(pb.Gender_value[string(res.Gender)]),
		Dob:       timestamppb.New(res.DateOfBirth),
		CreatedAt: timestamppb.New(res.CreatedAt),
		UpdatedAt: timestamppb.New(res.UpdatedAt),
	}, nil
}

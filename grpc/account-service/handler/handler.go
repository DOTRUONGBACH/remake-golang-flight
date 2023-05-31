package handler

import (
	"context"
	"jet/grpc/account-service/repository"
	"jet/pb"
	"log"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type AccountService struct {
	AccountRepository repository.AccountRepository
	pb.UnimplementedAccountServiceServer
}

func NewAccountHander(AccountRepository repository.AccountRepository) (*AccountService, error) {
	return &AccountService{
		AccountRepository: AccountRepository,
	}, nil
}

func (as AccountService) SignUp(ctx context.Context, model *pb.SingupRequest) (*pb.Account, error) {
	res, err := as.AccountRepository.CreateAccount(ctx, model)

	if err != nil {
		log.Println("Created failed: ", err)
		return nil, err
	}
	return &pb.Account{Id: res.ID.String(),
		Email:     res.Email,
		Password:  res.Password,
		Role:      pb.Role(pb.Role_value[res.Role.String()]),
		AccStatus: pb.AccountStatus(pb.AccountStatus_value[string(res.Status)]),
		CreatedAt: timestamppb.New(res.CreatedAt),
		UpdateAt:  timestamppb.New(res.UpdatedAt)}, nil
}

package handler

import (
	"context"
	"jet/grpc/account-service/repository"
	"jet/internal/jwt"
	"jet/internal/util"
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

func (as AccountService) LogIn(ctx context.Context, model *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := as.AccountRepository.GetAccountByEmail(ctx, &pb.GetAccountByEmailRequest{Email: model.Email})
	if err != nil {
		log.Print("acc does not exist: ", err)
		return &pb.LoginResponse{Token: "",
			Status: false}, err

	}

	if util.CheckPasswordHash(model.Password, user.Password) && user.Status == "active" {
		token, err := jwt.GenerateToken(user.Email)
		if err != nil {
			return &pb.LoginResponse{Token: "", Status: false}, err
		}
		return &pb.LoginResponse{Token: token, Status: true}, nil
	}

	return &pb.LoginResponse{Token: "", Status: false}, err
}

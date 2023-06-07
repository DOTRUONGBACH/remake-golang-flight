package services

import (
	"context"
	"jet/ent"
	"jet/ent/account"
	"jet/internal/util"
	"jet/pb"
	"log"

	"github.com/google/uuid"
)

type AccountServiceClient interface {
	Signup(ctx context.Context, input ent.NewAccountInput) (*ent.Account, error)
	Login(ctx context.Context, input ent.Login) (ent.AccountLoginResponse, error)
}

type AccountHandler struct {
	AccountClient  pb.AccountServiceClient
	CustomerClient pb.CustomerServiceClient
}

func NewAccountHandler(AccountClient pb.AccountServiceClient, CustomerClient pb.CustomerServiceClient) AccountServiceClient {
	return &AccountHandler{
		AccountClient:  AccountClient,
		CustomerClient: CustomerClient,
	}
}

func (a AccountHandler) Login(ctx context.Context, input ent.Login) (ent.AccountLoginResponse, error) {
	res, err := a.AccountClient.Login(ctx, &pb.LoginRequest{Email: input.Email, Password: input.Password})
	if err != nil || !res.Status {
		log.Println("Login fail:", err)
		return ent.AccountLoginResponse{Token: res.Token,
			Status: res.Status}, err
	}
	log.Println("Login success:")
	return ent.AccountLoginResponse{Token: res.Token,
		Status: res.Status}, nil

}

func (a AccountHandler) Signup(ctx context.Context, input ent.NewAccountInput) (*ent.Account, error) {
	hashPassword, _ := util.HashPassword(input.Password)
	newCustomer, err := a.CustomerClient.CreateCustomer(ctx, &pb.CreateCustomerRequest{
		Name:      input.Customer.Name,
		CitizenId: input.Customer.CitizenID,
		Phone:     input.Customer.Phone,
		Address:   input.Customer.Address,
		Gender:    pb.Gender(pb.Gender_value[string(input.Customer.Gender)]),
	})

	if err != nil {
		log.Fatalf("Created failed: %s", err)
	}

	newAccount, err := a.AccountClient.Signup(ctx, &pb.SingupRequest{
		Email:      input.Email,
		Password:   hashPassword,
		Role:       pb.Role(pb.Role_value[string(input.Role)]),
		AccOwnerId: newCustomer.Id,
	})

	if err != nil {
		return nil, err
	}

	return &ent.Account{
		ID:        uuid.MustParse(newAccount.Id),
		Email:     newAccount.Email,
		Role:      account.Role(newAccount.Role.String()),
		Status:    account.Status(newAccount.AccStatus),
		CreatedAt: newAccount.CreatedAt.AsTime(),
		UpdatedAt: newAccount.UpdateAt.AsTime(),
	}, nil
}

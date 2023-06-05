package repository

import (
	"context"
	"jet/ent"
	"jet/ent/account"
	"jet/pb"
	"log"
	"os"

	"github.com/google/uuid"
)

type AccountRepository interface {
	CreateAccount(ctx context.Context, model *pb.SingupRequest) (*ent.Account, error)
	CloseDB()
	GetAccountByEmail(ctx context.Context, model *pb.GetAccountByEmailRequest) (*ent.Account, error)
}

type PostgresDB struct {
	Client *ent.Client
}

func NewPostgresDB(connection_str string) (AccountRepository, error) {
	client, err := ent.Open("postgres", connection_str)
	if err != nil {
		return nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("Failed to creating db schema from the automation migration tool")
		os.Exit(1)
	}

	return &PostgresDB{Client: client}, nil
}

func (r *PostgresDB) CloseDB() {
	r.Client.Close()
}

func (r *PostgresDB) CreateAccount(ctx context.Context, model *pb.SingupRequest) (*ent.Account, error) {
	res, err := r.Client.Account.Create().SetEmail(model.Email).SetPassword(model.Password).SetRole(account.Role(model.Role)).SetStatus(account.StatusInactive).SetAccOwnerID(uuid.MustParse(model.AccOwnerId)).Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *PostgresDB) GetAccountByEmail(ctx context.Context, model *pb.GetAccountByEmailRequest) (*ent.Account, error) {
	res, err := r.Client.Account.Query().Where(account.Email(model.Email)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

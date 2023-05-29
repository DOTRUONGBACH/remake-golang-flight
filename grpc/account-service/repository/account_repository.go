package Repository

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
	Signup(ctx context.Context, model *pb.SingupRequest) (*ent.Account, *ent.Customer, error)
	Login(ctx context.Context, model *pb.LoginRequest) (*ent.Account, error)
	CloseDB()
}

type PostgresDB struct {
	Client *ent.Client
}

func NewPostgresDB(connection_str string) (*PostgresDB, error) {
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

func (r *PostgresDB) Signup(ctx context.Context, model *pb.SingupRequest) (*ent.Account, *ent.Customer, error) {
	res, err := r.Client.Account.Create().SetEmail(model.Email).SetPassword(model.Password).SetRole(account.Role(model.Role)).SetStatus(account.StatusInactive).SetAccOwnerID(model.AccOwnerId).Save(ctx)

}

func (r *PostgresDB) Login(ctx context.Context, model *pb.LoginRequest) (*ent.Account, error) {

}

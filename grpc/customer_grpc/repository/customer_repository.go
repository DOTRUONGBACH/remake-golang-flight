package repository

import (
	"context"
	"jet/ent"
	"jet/ent/customer"
	"jet/pb"
	"log"
	"os"
)

type CustomerRepository interface {
	CloseDB()

	CreateCustomer(ctx context.Context, model *pb.CreateCustomerRequest) (*ent.Customer, error)
}

type PostgresDB struct {
	Client *ent.Client
}

func NewPostgresDB(connection_str string) (CustomerRepository, error) {
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

func (r *PostgresDB) CreateCustomer(ctx context.Context, model *pb.CreateCustomerRequest) (*ent.Customer, error) {
	res, err := r.Client.Customer.Create().SetFullname(model.Name).SetCitizenID(model.CitizenId).SetPhone(model.Phone).SetAddress(model.Address).SetGender(customer.Gender(model.Gender.String())).SetDateOfBirth(model.Dob.AsTime()).Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

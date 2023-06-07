package services

import "jet/pb"

type CustomerServiceClient interface {
}

type CustomerHandler struct {
	CustomerClient pb.CustomerServiceClient
}

func NewCustomerHandler(CustomerClient pb.CustomerServiceClient) CustomerServiceClient {
	return &CustomerHandler{
		CustomerClient: CustomerClient,
	}
}


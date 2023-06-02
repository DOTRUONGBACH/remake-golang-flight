package resolver

import (
	"jet/ent"

	"github.com/go-playground/validator"
	"go.uber.org/zap"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client               *ent.Client
	validator            *validator.Validate
	validationTranslator ut.Translator
	logger               *zap.Logger
	customerService      services.CustomerServiceClient
	accountService       services.AccountServiceClient
}

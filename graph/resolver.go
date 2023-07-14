package graph

import "github.com/bmena06/crud_go/domain/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserUseCase usecase.UserUseCase
}

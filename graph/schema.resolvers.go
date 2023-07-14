package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"errors"

	"github.com/bmena06/crud_go/domain"
	"github.com/bmena06/crud_go/domain/entities"
	"github.com/bmena06/crud_go/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*entities.User, error) {
	if !domain.ValidateEmail(input.Email) {
		return nil, errors.New("el email no es válido")

	}

	if !domain.ValidateName(input.Name) {
		return nil, errors.New("el nombre no es valido")
	}

	user, err := r.UserUseCase.CreateUserUseCase(input)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.UpdateUserInput) (*entities.User, error) {
	if !domain.ValidateEmail(*input.Email) {
		return nil, errors.New("el email no es válido")
	}

	if !domain.ValidateName(*input.Name) {
		return nil, errors.New("el nombre no es valido")
	}

	user, err := r.UserUseCase.UpdateUserUseCase(id, input)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*model.DeleteUserResponse, error) {
	return r.UserUseCase.DeleteUserUseCase(id), nil
}

// SoftdeleteUser is the resolver for the softdeleteUser field.
func (r *mutationResolver) SoftdeleteUser(ctx context.Context, id string, input model.SoftdeleteUserInput) (*entities.User, error) {
	user, err := r.UserUseCase.SoftdeleteUserUseCase(id, input)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Getusers is the resolver for the getusers field.
func (r *queryResolver) Getusers(ctx context.Context, search *string, page *int, perpage *int) ([]*entities.User, error) {
	if search == nil {
		defaultSearch := ""
		search = &defaultSearch
	}

	if page == nil {
		defaultPage := 0
		page = &defaultPage
	}

	if perpage == nil {
		defaultPerpage := 0
		perpage = &defaultPerpage
	}

	users, err := r.UserUseCase.GetallUseCase(*search, *page, *perpage)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Getuser is the resolver for the getuser field.
func (r *queryResolver) Getuser(ctx context.Context, id string) (*entities.User, error) {
	users, err := r.UserUseCase.GetidUserCase(id)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
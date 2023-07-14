package usecase

import (
	"github.com/bmena06/crud_go/domain/entities"
	"github.com/bmena06/crud_go/domain/repositories"
	"github.com/bmena06/crud_go/graph/model"
)

type UserUseCase struct {
	Repository repositories.UserRepository
}

func (u *UserUseCase) CreateUserUseCase(user model.CreateUserInput) (*entities.User, error) {
	return u.Repository.CreateUser(user)
}

func (u *UserUseCase) GetallUseCase(searchQuery string, page int, perpage int) ([]*entities.User, error) {
	return u.Repository.Getall(searchQuery, page, perpage)
}

func (u *UserUseCase) SoftdeleteUserUseCase(id string, user model.SoftdeleteUserInput) (*entities.User, error) {
	return u.Repository.SoftdeleteUser(id, user)
}

func (u *UserUseCase) DeleteUserUseCase(id string) *model.DeleteUserResponse {
	return u.Repository.DeleteUser(id)
}

func (u *UserUseCase) UpdateUserUseCase(id string, user model.UpdateUserInput) (*entities.User, error) {
	return u.Repository.UpdateUser(id, user)
}

func (u *UserUseCase) GetidUserCase(id string) (*entities.User, error) {
	return u.Repository.Getid(id)
}

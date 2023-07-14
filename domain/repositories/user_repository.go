package repositories

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/bmena06/crud_go/domain/entities"
	"github.com/bmena06/crud_go/graph/model"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	Client *mongo.Database
}

// GET BY ID
func (r *UserRepository) Getid(id string) (*entities.User, error) {
	userCollec := r.Client.Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	var user entities.User
	err := userCollec.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("el usuario no existe")
		}
		return nil, err
	}

	if user.Deleted {
		return nil, errors.New("el usuario está eliminado")
	}

	return &user, nil
}

// GET ALL USERS
func (r *UserRepository) Getall(searchQuery string, page int, perpage int) ([]*entities.User, error) {
	userCollec := r.Client.Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var users []*entities.User
	filter := bson.M{
		"deleted": false, // Agrega la condición para excluir los registros con deleted: true
	}
	if searchQuery != "" {
		filter["name"] = bson.M{"$regex": searchQuery, "$options": "i"}
	}

	options := options.Find().SetSort(bson.D{{Key: "name", Value: 1}}).SetSkip(int64(page)).SetLimit(int64(perpage)) // Orden ascendente por el campo "name"
	cursor, err := userCollec.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

// CREATE USERS
func (r *UserRepository) CreateUser(userInfo model.CreateUserInput) (*entities.User, error) {
	userCollec := r.Client.Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Verificar si el correo electrónico ya existe en la base de datos
	emailExists, err := r.validateRepeated("email", userInfo.Email)
	if err != nil {
		return nil, err
	}
	if emailExists {
		return nil, errors.New("el correo electrónico ya está en uso")
	}

	// Verificar si el número de teléfono ya existe en la base de datos
	phoneExists, err := r.validateRepeated("phone", userInfo.Phone)
	if err != nil {
		return nil, err
	}
	if phoneExists {
		return nil, errors.New("el número de teléfono ya está en uso")
	}

	id := uuid.New().String()                                                                                                                        // Genera un ID único
	_, err = userCollec.InsertOne(ctx, bson.M{"_id": id, "name": userInfo.Name, "phone": userInfo.Phone, "email": userInfo.Email, "deleted": false}) // Establecer deleted en false
	if err != nil {
		return nil, err
	}

	returnUserListing := entities.User{ID: id, Name: userInfo.Name, Phone: userInfo.Phone, Email: userInfo.Email, Deleted: false} // Establecer Deleted en false
	return &returnUserListing, nil
}

func (r *UserRepository) validateRepeated(fieldName string, value string) (bool, error) {
	userCollec := r.Client.Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{fieldName: value}
	count, err := userCollec.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// UPDATE USER
func (r *UserRepository) UpdateUser(id string, userInfo model.UpdateUserInput) (*entities.User, error) {
	userCollec := r.Client.Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Verificar si el correo electrónico ya existe en la base de datos
	if userInfo.Email != nil {
		emailExists, err := r.validateRepeated("email", *userInfo.Email)
		if err != nil {
			return nil, err
		}
		if emailExists {
			return nil, errors.New("el correo electrónico ya está en uso")
		}
	}

	// Verificar si el número de teléfono ya existe en la base de datos
	if userInfo.Phone != nil {
		phoneExists, err := r.validateRepeated("phone", *userInfo.Phone)
		if err != nil {
			return nil, err
		}
		if phoneExists {
			return nil, errors.New("el número de teléfono ya está en uso")
		}
	}

	updateUserInfo := bson.M{}

	if userInfo.Name != nil {
		updateUserInfo["name"] = *userInfo.Name
	}
	if userInfo.Phone != nil {
		updateUserInfo["phone"] = *userInfo.Phone
	}
	if userInfo.Email != nil {
		updateUserInfo["email"] = *userInfo.Email
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": updateUserInfo}

	options := options.FindOneAndUpdate().SetReturnDocument(options.After)

	result := userCollec.FindOneAndUpdate(ctx, filter, update, options)
	if result.Err() != nil {
		return nil, result.Err()
	}

	var user entities.User
	if err := result.Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

// DELETE USER
func (r *UserRepository) DeleteUser(id string) *model.DeleteUserResponse {
	userCollec := r.Client.Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	_, err := userCollec.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	return &model.DeleteUserResponse{DeletedUserID: id}
}

// SOFT DELETE USER
func (r *UserRepository) SoftdeleteUser(id string, userInfo model.SoftdeleteUserInput) (*entities.User, error) {
	userCollec := r.Client.Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	toDeleteUserInfo := bson.M{}

	if userInfo.Deleted {
		toDeleteUserInfo["deleted"] = userInfo.Deleted
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": toDeleteUserInfo}

	options := options.FindOneAndUpdate().SetReturnDocument(options.After)

	result := userCollec.FindOneAndUpdate(ctx, filter, update, options)
	if result.Err() != nil {
		return nil, result.Err()
	}

	var user entities.User
	if err := result.Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

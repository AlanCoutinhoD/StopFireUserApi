package repositories

import (
	"context"

	"hex_go/src/users/domain/entities"
)

// UserRepository define las operaciones que se pueden realizar con la entidad User
type UserRepository interface {
	Create(ctx context.Context, user *entities.User) (*entities.User, error)
	FindByID(ctx context.Context, id int) (*entities.User, error)
	FindByUsername(ctx context.Context, username string) (*entities.User, error)
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	Update(ctx context.Context, user *entities.User) error
	Delete(ctx context.Context, id int) error
}
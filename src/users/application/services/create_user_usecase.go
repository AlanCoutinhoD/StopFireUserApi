package services

import (
	"context"
	"errors"

	"hex_go/src/users/domain/entities"
	"hex_go/src/users/domain/repositories"
	"golang.org/x/crypto/bcrypt"
)

// CreateUserUseCase implementa el caso de uso para crear un usuario
type CreateUserUseCase struct {
	userRepository repositories.UserRepository
}

// NewCreateUserUseCase crea una nueva instancia de CreateUserUseCase
func NewCreateUserUseCase(userRepo repositories.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository: userRepo,
	}
}

// Execute ejecuta el caso de uso
func (uc *CreateUserUseCase) Execute(ctx context.Context, username, password, email string) (*entities.User, error) {
	// Verificar si el usuario ya existe
	existingUser, _ := uc.userRepository.FindByUsername(ctx, username)
	if existingUser != nil {
		return nil, errors.New("username already exists")
	}

	// Verificar si el email ya existe
	existingEmail, _ := uc.userRepository.FindByEmail(ctx, email)
	if existingEmail != nil {
		return nil, errors.New("email already exists")
	}

	// Encriptar la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Crear el usuario
	user := entities.NewUser(username, string(hashedPassword), email)
	
	// Guardar el usuario
	createdUser, err := uc.userRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
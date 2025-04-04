package services

import (
	"context"
	"errors"

	"hex_go/src/esp32/domain/entities"
	"hex_go/src/esp32/domain/repositories"
	userRepo "hex_go/src/users/domain/repositories"
)

// GetUserESP32sUseCase implementa el caso de uso para obtener todos los ESP32 de un usuario
type GetUserESP32sUseCase struct {
	esp32Repository repositories.ESP32Repository
	userRepository  userRepo.UserRepository
}

// NewGetUserESP32sUseCase crea una nueva instancia de GetUserESP32sUseCase
func NewGetUserESP32sUseCase(esp32Repo repositories.ESP32Repository, userRepo userRepo.UserRepository) *GetUserESP32sUseCase {
	return &GetUserESP32sUseCase{
		esp32Repository: esp32Repo,
		userRepository:  userRepo,
	}
}

// Execute ejecuta el caso de uso
func (uc *GetUserESP32sUseCase) Execute(ctx context.Context, userID int) ([]*entities.ESP32, error) {
	// Verificar si el usuario existe
	user, err := uc.userRepository.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// Obtener todos los ESP32 del usuario
	return uc.esp32Repository.FindByUserID(ctx, userID)
}
package services

import (
	"context"
	"errors"

	"hex_go/src/esp32/domain/repositories"
	userRepo "hex_go/src/users/domain/repositories"
)

// AssignESP32UseCase implementa el caso de uso para asignar un ESP32 a un usuario
type AssignESP32UseCase struct {
	esp32Repository repositories.ESP32Repository
	userRepository  userRepo.UserRepository
}

// NewAssignESP32UseCase crea una nueva instancia de AssignESP32UseCase
func NewAssignESP32UseCase(esp32Repo repositories.ESP32Repository, userRepo userRepo.UserRepository) *AssignESP32UseCase {
	return &AssignESP32UseCase{
		esp32Repository: esp32Repo,
		userRepository:  userRepo,
	}
}

// Execute ejecuta el caso de uso
func (uc *AssignESP32UseCase) Execute(ctx context.Context, esp32ID, userID int) error {
	// Verificar si el ESP32 existe
	esp32, err := uc.esp32Repository.FindByID(ctx, esp32ID)
	if err != nil {
		return err
	}
	if esp32 == nil {
		return errors.New("ESP32 not found")
	}

	// Verificar si el usuario existe
	user, err := uc.userRepository.FindByID(ctx, userID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}

	// Verificar si el ESP32 ya está asignado a otro usuario
	if esp32.UserID != nil && *esp32.UserID != userID {
		return errors.New("ESP32 already assigned to another user")
	}

	// Asignar el ESP32 al usuario
	return uc.esp32Repository.AssignToUser(ctx, esp32ID, userID)
}
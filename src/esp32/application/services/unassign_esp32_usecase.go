package services

import (
	"context"
	"errors"

	"hex_go/src/esp32/domain/repositories"
)

// UnassignESP32UseCase implementa el caso de uso para desasignar un ESP32 de un usuario
type UnassignESP32UseCase struct {
	esp32Repository repositories.ESP32Repository
}

// NewUnassignESP32UseCase crea una nueva instancia de UnassignESP32UseCase
func NewUnassignESP32UseCase(esp32Repo repositories.ESP32Repository) *UnassignESP32UseCase {
	return &UnassignESP32UseCase{
		esp32Repository: esp32Repo,
	}
}

// Execute ejecuta el caso de uso
func (uc *UnassignESP32UseCase) Execute(ctx context.Context, esp32ID int) error {
	// Verificar si el ESP32 existe
	esp32, err := uc.esp32Repository.FindByID(ctx, esp32ID)
	if err != nil {
		return err
	}
	if esp32 == nil {
		return errors.New("ESP32 not found")
	}

	// Verificar si el ESP32 está asignado a algún usuario
	if esp32.UserID == nil {
		return errors.New("ESP32 is not assigned to any user")
	}

	// Desasignar el ESP32
	return uc.esp32Repository.UnassignFromUser(ctx, esp32ID)
}
package repositories

import (
	"context"

	"hex_go/src/esp32/domain/entities"
)

// ESP32Repository define las operaciones que se pueden realizar con la entidad ESP32
type ESP32Repository interface {
	Create(ctx context.Context, esp32 *entities.ESP32) (*entities.ESP32, error)
	FindByID(ctx context.Context, id int) (*entities.ESP32, error)
	FindByNumeroSerie(ctx context.Context, numeroSerie string) (*entities.ESP32, error)
	FindByUserID(ctx context.Context, userID int) ([]*entities.ESP32, error)
	FindUnassigned(ctx context.Context) ([]*entities.ESP32, error)
	Update(ctx context.Context, esp32 *entities.ESP32) error
	Delete(ctx context.Context, id int) error
	AssignToUser(ctx context.Context, esp32ID, userID int) error
	UnassignFromUser(ctx context.Context, esp32ID int) error
}
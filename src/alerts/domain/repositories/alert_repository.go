package repositories

import (
	"context"
	"hex_go/src/alerts/domain/entities"
)

// AlertRepository defines operations for alert data
type AlertRepository interface {
	GetAlertsByUserID(ctx context.Context, userID int) ([]*entities.Alert, error)
	GetAlertsByESP32ID(ctx context.Context, esp32ID int) ([]*entities.Alert, error)
	GetAlertsByESP32NumeroSerie(ctx context.Context, numeroSerie string) ([]*entities.Alert, error)
}
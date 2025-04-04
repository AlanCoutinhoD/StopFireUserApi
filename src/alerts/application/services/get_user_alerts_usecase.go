package services

import (
	"context"
	"hex_go/src/alerts/domain/entities"
	"hex_go/src/alerts/domain/repositories"
)

// GetUserAlertsUseCase handles getting alerts for a user
type GetUserAlertsUseCase struct {
	alertRepository repositories.AlertRepository
}

// NewGetUserAlertsUseCase creates a new instance of GetUserAlertsUseCase
func NewGetUserAlertsUseCase(alertRepository repositories.AlertRepository) *GetUserAlertsUseCase {
	return &GetUserAlertsUseCase{
		alertRepository: alertRepository,
	}
}

// Execute gets all alerts for a user
func (uc *GetUserAlertsUseCase) Execute(ctx context.Context, userID int) ([]*entities.Alert, error) {
	return uc.alertRepository.GetAlertsByUserID(ctx, userID)
}
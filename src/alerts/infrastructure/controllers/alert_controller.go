package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hex_go/src/alerts/application/services"
)

// AlertController handles HTTP requests for alerts
type AlertController struct {
	getUserAlertsUseCase *services.GetUserAlertsUseCase
}

// NewAlertController creates a new instance of AlertController
func NewAlertController(getUserAlertsUseCase *services.GetUserAlertsUseCase) *AlertController {
	return &AlertController{
		getUserAlertsUseCase: getUserAlertsUseCase,
	}
}

// GetUserAlerts handles the HTTP request to get all alerts for a user
func (c *AlertController) GetUserAlerts(ctx *gin.Context) {
	// Get the user ID from the context (set by the authentication middleware)
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	alerts, err := c.getUserAlertsUseCase.Execute(ctx, userID.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, alerts)
}

// SetupRoutes configures the routes for the alert controller
func (c *AlertController) SetupRoutes(router *gin.Engine, authMiddleware gin.HandlerFunc) {
	api := router.Group("/api")
	{
		alerts := api.Group("/alerts")
		{
			// Protected routes (require authentication)
			protected := alerts.Group("")
			protected.Use(authMiddleware)
			{
				protected.GET("/user", c.GetUserAlerts)
			}
		}
	}
}
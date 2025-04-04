package infrastructure

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"hex_go/src/alerts/application/services"
	"hex_go/src/alerts/infrastructure/controllers"
	"hex_go/src/alerts/infrastructure/repositories"
	"hex_go/src/middleware" 
)

// Init initializes the alerts module
func Init(router *gin.Engine, db *sql.DB) {
	log.Println("Initializing alerts module...")

	// Initialize repositories
	alertRepo := repositories.NewMySQLAlertRepository(db)

	// Initialize use cases
	getUserAlertsUseCase := services.NewGetUserAlertsUseCase(alertRepo)

	// Initialize controllers
	alertController := controllers.NewAlertController(getUserAlertsUseCase)

	// Get auth middleware
	authMiddleware := middleware.AuthMiddleware()

	// Setup routes
	alertController.SetupRoutes(router, authMiddleware)
}
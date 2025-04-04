package infrastructure

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"hex_go/src/config"
	"hex_go/src/users/application/services"
	"hex_go/src/users/infrastructure/controllers"
	"hex_go/src/users/infrastructure/repositories"
)

// Init inicializa la infraestructura de usuarios
func Init(router *gin.Engine) {
	// Inicializar base de datos
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Crear tabla de usuarios si no existe
	createUsersTable(db)

	// Inicializar repositorios
	userRepo := repositories.NewMySQLUserRepository(db)

	// Inicializar casos de uso
	createUserUseCase := services.NewCreateUserUseCase(userRepo)
	loginUserUseCase := services.NewLoginUserUseCase(userRepo)

	// Inicializar controladores
	userController := controllers.NewUserController(createUserUseCase, loginUserUseCase)

	// Configurar rutas
	userController.SetupRoutes(router)
}

// createUsersTable crea la tabla de usuarios si no existe
func createUsersTable(db *sql.DB) {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			username VARCHAR(255) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL UNIQUE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`
	
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}
}
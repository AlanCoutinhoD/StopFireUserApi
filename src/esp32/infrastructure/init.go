package infrastructure

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"hex_go/src/esp32/application/services"
	"hex_go/src/esp32/infrastructure/controllers"
	"hex_go/src/esp32/infrastructure/repositories"
	"hex_go/src/middleware"
	userRepo "hex_go/src/users/infrastructure/repositories"
)

// Init inicializa la infraestructura de ESP32
func Init(router *gin.Engine, db *sql.DB) {
	// Crear tabla de ESP32 si no existe
	createESP32Table(db)

	// Inicializar repositorios
	esp32Repo := repositories.NewMySQLESP32Repository(db)
	userRepository := userRepo.NewMySQLUserRepository(db)

	// Inicializar casos de uso
	assignESP32UseCase := services.NewAssignESP32UseCase(esp32Repo, userRepository)
	unassignESP32UseCase := services.NewUnassignESP32UseCase(esp32Repo)
	getUserESP32sUseCase := services.NewGetUserESP32sUseCase(esp32Repo, userRepository)

	// Inicializar controladores
	esp32Controller := controllers.NewESP32Controller(
		assignESP32UseCase,
		unassignESP32UseCase,
		getUserESP32sUseCase,
		esp32Repo,
	)

	// Configurar rutas
	authMiddleware := middleware.AuthMiddleware()
	esp32Controller.SetupRoutes(router, authMiddleware)
}

// createESP32Table crea la tabla de ESP32 si no existe
func createESP32Table(db *sql.DB) {
	// First check if the table exists
	var tableName string
	err := db.QueryRow("SHOW TABLES LIKE 'esp32'").Scan(&tableName)
	if err == nil {
		// Table already exists, no need to create it
		log.Println("ESP32 table already exists, skipping creation")
		return
	}

	// Table doesn't exist, create it
	query := `
		CREATE TABLE IF NOT EXISTS esp32 (
			idESP32 INT AUTO_INCREMENT PRIMARY KEY,
			idKY_026 INT,
			idMQ_2 INT,
			idMQ_135 INT,
			idDHT_22 INT,
			numero_serie VARCHAR(255) NOT NULL UNIQUE,
			idUser INT,
			FOREIGN KEY (idUser) REFERENCES users(id) ON DELETE SET NULL
		)
	`
	
	_, err = db.Exec(query)
	if err != nil {
		log.Printf("Warning: Failed to create ESP32 table: %v", err)
		// Don't fatally exit, just log the warning
	} else {
		log.Println("ESP32 table created successfully")
	}
}
package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"hex_go/src/config"
	"hex_go/src/esp32/infrastructure"
	userInfrastructure "hex_go/src/users/infrastructure"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// Inicializar base de datos
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	router := gin.Default()

	// Configuración de CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

	router.Use(cors.New(config))

	// Add a simple root route to check if the server is running
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API is running",
		})
	})

	// Inicializar infraestructura de usuarios
	userInfrastructure.Init(router)

	// Inicializar infraestructura de ESP32
	infrastructure.Init(router, db)

	// Iniciar el servidor
	log.Println("Server running on port 8080")
	router.Run(":8080")
}
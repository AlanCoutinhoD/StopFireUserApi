package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// InitDB inicializa la conexión a la base de datos
func InitDB() (*sql.DB, error) {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using default values")
	}

	// Obtener configuración de la base de datos
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "3306")
	dbUser := getEnv("DB_USER", "root")
	dbPassword := getEnv("DB_PASSWORD", "")  // Make sure empty password is handled correctly
	dbName := getEnv("DB_NAME", "stopfire")

	// Construir DSN (Data Source Name)
	var dsn string
	if dbPassword == "" {
		// DSN without password
		dsn = fmt.Sprintf("%s@tcp(%s:%s)/%s?parseTime=true", 
			dbUser, dbHost, dbPort, dbName)
	} else {
		// DSN with password
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", 
			dbUser, dbPassword, dbHost, dbPort, dbName)
	}

	log.Printf("Connecting to database: %s@tcp(%s:%s)/%s", dbUser, dbHost, dbPort, dbName)
	
	// Abrir conexión
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Verificar conexión
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to database successfully")
	return db, nil
}

// getEnv obtiene una variable de entorno o devuelve un valor por defecto
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
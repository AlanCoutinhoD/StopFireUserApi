package entities

import "time"

// AlertType represents the type of sensor that triggered the alert
type AlertType string

const (
	AlertTypeKY026 AlertType = "KY_026"
	AlertTypeMQ2   AlertType = "MQ_2"
	AlertTypeMQ135 AlertType = "MQ_135"
	AlertTypeDHT22 AlertType = "DHT_22"
)

// Alert represents a sensor alert
type Alert struct {
	ID              int       `json:"id"`
	ESP32ID         int       `json:"esp32_id"`
	ESP32NumeroSerie string    `json:"esp32_numero_serie"`
	SensorID        int       `json:"sensor_id"`
	SensorType      AlertType `json:"sensor_type"`
	Estado          int       `json:"estado"`
	FechaActivacion string    `json:"fecha_activacion"`
	CreatedAt       time.Time `json:"created_at"`
}
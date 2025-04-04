package entities

import (
	"time"
)

// ESP32 representa la entidad de dominio para un dispositivo ESP32
type ESP32 struct {
	ID          int       `json:"id"`
	IDKY026     int       `json:"id_ky_026"`
	IDMQ2       int       `json:"id_mq_2"`
	IDMQ135     int       `json:"id_mq_135"`
	IDDHT22     int       `json:"id_dht_22"`
	NumeroSerie string    `json:"numero_serie"`
	UserID      *int      `json:"user_id"` // Puede ser nulo si no está asignado
	CreatedAt   time.Time `json:"created_at"`
}

// NewESP32 crea una nueva instancia de ESP32
func NewESP32(idKY026, idMQ2, idMQ135, idDHT22 int, numeroSerie string) *ESP32 {
	return &ESP32{
		IDKY026:     idKY026,
		IDMQ2:       idMQ2,
		IDMQ135:     idMQ135,
		IDDHT22:     idDHT22,
		NumeroSerie: numeroSerie,
		CreatedAt:   time.Now(),
	}
}

// AssignToUser asigna el ESP32 a un usuario
func (e *ESP32) AssignToUser(userID int) {
	e.UserID = &userID
}

// UnassignFromUser desasigna el ESP32 de cualquier usuario
func (e *ESP32) UnassignFromUser() {
	e.UserID = nil
}
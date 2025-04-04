package repositories

import (
	"context"
	"database/sql"
	"time"
	
	"hex_go/src/alerts/domain/entities"
	"hex_go/src/alerts/domain/repositories"
)

// MySQLAlertRepository implements AlertRepository using MySQL
type MySQLAlertRepository struct {
	db *sql.DB
}

// NewMySQLAlertRepository creates a new instance of MySQLAlertRepository
func NewMySQLAlertRepository(db *sql.DB) repositories.AlertRepository {
	return &MySQLAlertRepository{
		db: db,
	}
}

// GetAlertsByUserID retrieves all alerts for a specific user
func (r *MySQLAlertRepository) GetAlertsByUserID(ctx context.Context, userID int) ([]*entities.Alert, error) {
	query := `
		SELECT 
			s.idKY_026 as sensor_id, 
			'KY_026' as sensor_type, 
			s.estado, 
			s.fecha_activacion, 
			e.idESP32, 
			e.numero_serie
		FROM KY_026 s
		JOIN ESP32 e ON s.idKY_026 = e.idKY_026
		WHERE e.idUser = ? AND s.estado = 1
		UNION
		SELECT 
			s.idMQ_2 as sensor_id, 
			'MQ_2' as sensor_type, 
			s.estado, 
			s.fecha_activacion, 
			e.idESP32, 
			e.numero_serie
		FROM MQ_2 s
		JOIN ESP32 e ON s.idMQ_2 = e.idMQ_2
		WHERE e.idUser = ? AND s.estado = 1
		UNION
		SELECT 
			s.idMQ_135 as sensor_id, 
			'MQ_135' as sensor_type, 
			s.estado, 
			s.fecha_activacion, 
			e.idESP32, 
			e.numero_serie
		FROM MQ_135 s
		JOIN ESP32 e ON s.idMQ_135 = e.idMQ_135
		WHERE e.idUser = ? AND s.estado = 1
		UNION
		SELECT 
			s.idDHT_22 as sensor_id, 
			'DHT_22' as sensor_type, 
			s.estado, 
			s.fecha_activacion, 
			e.idESP32, 
			e.numero_serie
		FROM DHT_22 s
		JOIN ESP32 e ON s.idDHT_22 = e.idDHT_22
		WHERE e.idUser = ? AND s.estado = 1
		ORDER BY fecha_activacion DESC
	`

	rows, err := r.db.QueryContext(ctx, query, userID, userID, userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alerts []*entities.Alert

	for rows.Next() {
		var alert entities.Alert
		var sensorType string

		err := rows.Scan(
			&alert.SensorID,
			&sensorType,
			&alert.Estado,
			&alert.FechaActivacion,
			&alert.ESP32ID,
			&alert.ESP32NumeroSerie,
		)

		if err != nil {
			return nil, err
		}

		alert.SensorType = entities.AlertType(sensorType)
		alert.CreatedAt = time.Now() // Since we don't have this in the database

		alerts = append(alerts, &alert)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return alerts, nil
}

// GetAlertsByESP32ID retrieves all alerts for a specific ESP32 by ID
func (r *MySQLAlertRepository) GetAlertsByESP32ID(ctx context.Context, esp32ID int) ([]*entities.Alert, error) {
	query := `
		SELECT 
			s.idKY_026 as sensor_id, 
			'KY_026' as sensor_type, 
			s.estado, 
			s.fecha_activacion, 
			e.idESP32, 
			e.numero_serie
		FROM KY_026 s
		JOIN ESP32 e ON s.idKY_026 = e.idKY_026
		WHERE e.idESP32 = ? AND s.estado = 1
		UNION
		SELECT 
			s.idMQ_2 as sensor_id, 
			'MQ_2' as sensor_type, 
			s.estado, 
			s.fecha_activacion, 
			e.idESP32, 
			e.numero_serie
		FROM MQ_2 s
		JOIN ESP32 e ON s.idMQ_2 = e.idMQ_2
		WHERE e.idESP32 = ? AND s.estado = 1
		UNION
		SELECT 
			s.idMQ_135 as sensor_id, 
			'MQ_135' as sensor_type, 
			s.estado, 
			s.fecha_activacion, 
			e.idESP32, 
			e.numero_serie
		FROM MQ_135 s
		JOIN ESP32 e ON s.idMQ_135 = e.idMQ_135
		WHERE e.idESP32 = ? AND s.estado = 1
		UNION
		SELECT 
			s.idDHT_22 as sensor_id, 
			'DHT_22' as sensor_type, 
			s.estado, 
			s.fecha_activacion, 
			e.idESP32, 
			e.numero_serie
		FROM DHT_22 s
		JOIN ESP32 e ON s.idDHT_22 = e.idDHT_22
		WHERE e.idESP32 = ? AND s.estado = 1
		ORDER BY fecha_activacion DESC
	`

	rows, err := r.db.QueryContext(ctx, query, esp32ID, esp32ID, esp32ID, esp32ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alerts []*entities.Alert

	for rows.Next() {
		var alert entities.Alert
		var sensorType string

		err := rows.Scan(
			&alert.SensorID,
			&sensorType,
			&alert.Estado,
			&alert.FechaActivacion,
			&alert.ESP32ID,
			&alert.ESP32NumeroSerie,
		)

		if err != nil {
			return nil, err
		}

		alert.SensorType = entities.AlertType(sensorType)
		alert.CreatedAt = time.Now() // Since we don't have this in the database

		alerts = append(alerts, &alert)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return alerts, nil
}

// GetAlertsByESP32NumeroSerie retrieves all alerts for a specific ESP32 by serial number
func (r *MySQLAlertRepository) GetAlertsByESP32NumeroSerie(ctx context.Context, numeroSerie string) ([]*entities.Alert, error) {
	query := `
		SELECT 
			s.idKY_026 as sensor_id, 
			'KY_026' as sensor_type, 
			s.estado, 
			s.fecha_activacion, 
			e.idESP32, 
			e.numero_serie
		FROM KY_026 s
		JOIN ESP32 e ON s.idKY_026 = e.idKY_026
		WHERE e.numero_serie = ? AND s.estado = 1
		UNION
		SELECT 
			s.idMQ_2 as sensor_id, 
			'MQ_2' as sensor_type, 
			s.estado, 
			s.fecha_activacion, 
			e.idESP32, 
			e.numero_serie
		FROM MQ_2 s
		JOIN ESP32 e ON s.idMQ_2 = e.idMQ_2
		WHERE e.numero_serie = ? AND s.estado = 1
		UNION
		SELECT 
			s.idMQ_135 as sensor_id, 
			'MQ_135' as sensor_type, 
			s.estado, 
			s.fecha_activacion, 
			e.idESP32, 
			e.numero_serie
		FROM MQ_135 s
		JOIN ESP32 e ON s.idMQ_135 = e.idMQ_135
		WHERE e.numero_serie = ? AND s.estado = 1
		UNION
		SELECT 
			s.idDHT_22 as sensor_id, 
			'DHT_22' as sensor_type, 
			s.estado, 
			s.fecha_activacion, 
			e.idESP32, 
			e.numero_serie
		FROM DHT_22 s
		JOIN ESP32 e ON s.idDHT_22 = e.idDHT_22
		WHERE e.numero_serie = ? AND s.estado = 1
		ORDER BY fecha_activacion DESC
	`

	rows, err := r.db.QueryContext(ctx, query, numeroSerie, numeroSerie, numeroSerie, numeroSerie)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alerts []*entities.Alert

	for rows.Next() {
		var alert entities.Alert
		var sensorType string

		err := rows.Scan(
			&alert.SensorID,
			&sensorType,
			&alert.Estado,
			&alert.FechaActivacion,
			&alert.ESP32ID,
			&alert.ESP32NumeroSerie,
		)

		if err != nil {
			return nil, err
		}

		alert.SensorType = entities.AlertType(sensorType)
		alert.CreatedAt = time.Now() // Since we don't have this in the database

		alerts = append(alerts, &alert)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return alerts, nil
}
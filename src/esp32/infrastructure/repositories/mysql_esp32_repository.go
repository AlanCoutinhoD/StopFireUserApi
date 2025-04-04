package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"hex_go/src/esp32/domain/entities"
	"hex_go/src/esp32/domain/repositories"
)

// MySQLESP32Repository implementa ESP32Repository usando MySQL
type MySQLESP32Repository struct {
	db *sql.DB
}

// NewMySQLESP32Repository crea una nueva instancia de MySQLESP32Repository
func NewMySQLESP32Repository(db *sql.DB) repositories.ESP32Repository {
	return &MySQLESP32Repository{
		db: db,
	}
}

// Create inserta un nuevo ESP32 en la base de datos
func (r *MySQLESP32Repository) Create(ctx context.Context, esp32 *entities.ESP32) (*entities.ESP32, error) {
	query := `INSERT INTO esp32 (idKY_026, idMQ_2, idMQ_135, idDHT_22, numero_serie, idUser) 
              VALUES (?, ?, ?, ?, ?, ?)`
	
	var userID interface{}
	if esp32.UserID != nil {
		userID = *esp32.UserID
	} else {
		userID = nil
	}
	
	result, err := r.db.ExecContext(ctx, query, 
		esp32.IDKY026, esp32.IDMQ2, esp32.IDMQ135, esp32.IDDHT22, esp32.NumeroSerie, userID)
	if err != nil {
		return nil, err
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	
	esp32.ID = int(id)
	
	return esp32, nil
}

// FindByID busca un ESP32 por su ID
func (r *MySQLESP32Repository) FindByID(ctx context.Context, id int) (*entities.ESP32, error) {
	query := `SELECT idESP32, idKY_026, idMQ_2, idMQ_135, idDHT_22, numero_serie, idUser
              FROM esp32 WHERE idESP32 = ?`
	
	var esp32 entities.ESP32
	var userID sql.NullInt64
	var idKY026 sql.NullInt64
	var idMQ2 sql.NullInt64
	var idMQ135 sql.NullInt64
	var idDHT22 sql.NullInt64
	
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&esp32.ID,
		&idKY026,
		&idMQ2,
		&idMQ135,
		&idDHT22,
		&esp32.NumeroSerie,
		&userID,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("ESP32 not found")
		}
		return nil, err
	}
	
	// Convert nullable fields to int
	if idKY026.Valid {
		esp32.IDKY026 = int(idKY026.Int64)
	}
	if idMQ2.Valid {
		esp32.IDMQ2 = int(idMQ2.Int64)
	}
	if idMQ135.Valid {
		esp32.IDMQ135 = int(idMQ135.Int64)
	}
	if idDHT22.Valid {
		esp32.IDDHT22 = int(idDHT22.Int64)
	}
	
	if userID.Valid {
		userIDInt := int(userID.Int64)
		esp32.UserID = &userIDInt
	}
	
	// Set current time as created_at since it's not in the database
	esp32.CreatedAt = time.Now()
	
	return &esp32, nil
}

// FindByUserID busca todos los ESP32 asignados a un usuario
func (r *MySQLESP32Repository) FindByUserID(ctx context.Context, userID int) ([]*entities.ESP32, error) {
	query := `SELECT idESP32, idKY_026, idMQ_2, idMQ_135, idDHT_22, numero_serie, idUser
              FROM esp32 WHERE idUser = ?`
	
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var esp32s []*entities.ESP32
	
	for rows.Next() {
		var esp32 entities.ESP32
		var userIDSQL sql.NullInt64
		var idKY026 sql.NullInt64
		var idMQ2 sql.NullInt64
		var idMQ135 sql.NullInt64
		var idDHT22 sql.NullInt64
		
		err := rows.Scan(
			&esp32.ID,
			&idKY026,
			&idMQ2,
			&idMQ135,
			&idDHT22,
			&esp32.NumeroSerie,
			&userIDSQL,
		)
		
		if err != nil {
			return nil, err
		}
		
		// Convert nullable fields to int
		if idKY026.Valid {
			esp32.IDKY026 = int(idKY026.Int64)
		}
		if idMQ2.Valid {
			esp32.IDMQ2 = int(idMQ2.Int64)
		}
		if idMQ135.Valid {
			esp32.IDMQ135 = int(idMQ135.Int64)
		}
		if idDHT22.Valid {
			esp32.IDDHT22 = int(idDHT22.Int64)
		}
		
		if userIDSQL.Valid {
			userIDInt := int(userIDSQL.Int64)
			esp32.UserID = &userIDInt
		}
		
		// Set current time as created_at since it's not in the database
		esp32.CreatedAt = time.Now()
		
		esp32s = append(esp32s, &esp32)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return esp32s, nil
}

// FindUnassigned busca todos los ESP32 no asignados a ningún usuario
func (r *MySQLESP32Repository) FindUnassigned(ctx context.Context) ([]*entities.ESP32, error) {
	query := `SELECT idESP32, idKY_026, idMQ_2, idMQ_135, idDHT_22, numero_serie, idUser
              FROM esp32 WHERE idUser IS NULL`
	
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var esp32s []*entities.ESP32
	
	for rows.Next() {
		var esp32 entities.ESP32
		var userIDSQL sql.NullInt64
		var idKY026 sql.NullInt64
		var idMQ2 sql.NullInt64
		var idMQ135 sql.NullInt64
		var idDHT22 sql.NullInt64
		
		err := rows.Scan(
			&esp32.ID,
			&idKY026,
			&idMQ2,
			&idMQ135,
			&idDHT22,
			&esp32.NumeroSerie,
			&userIDSQL,
		)
		
		if err != nil {
			return nil, err
		}
		
		// Convert nullable fields to int
		if idKY026.Valid {
			esp32.IDKY026 = int(idKY026.Int64)
		}
		if idMQ2.Valid {
			esp32.IDMQ2 = int(idMQ2.Int64)
		}
		if idMQ135.Valid {
			esp32.IDMQ135 = int(idMQ135.Int64)
		}
		if idDHT22.Valid {
			esp32.IDDHT22 = int(idDHT22.Int64)
		}
		
		if userIDSQL.Valid {
			userIDInt := int(userIDSQL.Int64)
			esp32.UserID = &userIDInt
		}
		
		// Set current time as created_at since it's not in the database
		esp32.CreatedAt = time.Now()
		
		esp32s = append(esp32s, &esp32)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return esp32s, nil
}

// Update actualiza un ESP32 existente
func (r *MySQLESP32Repository) Update(ctx context.Context, esp32 *entities.ESP32) error {
	query := `UPDATE esp32 SET idKY_026 = ?, idMQ_2 = ?, idMQ_135 = ?, idDHT_22 = ?, 
              numero_serie = ?, idUser = ? WHERE idESP32 = ?`
	
	var userID interface{}
	if esp32.UserID != nil {
		userID = *esp32.UserID
	} else {
		userID = nil
	}
	
	_, err := r.db.ExecContext(ctx, query, 
		esp32.IDKY026, esp32.IDMQ2, esp32.IDMQ135, esp32.IDDHT22, esp32.NumeroSerie, userID, esp32.ID)
	return err
}

// Delete elimina un ESP32 por su ID
func (r *MySQLESP32Repository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM esp32 WHERE idESP32 = ?`
	
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// AssignToUser asigna un ESP32 a un usuario
func (r *MySQLESP32Repository) AssignToUser(ctx context.Context, esp32ID, userID int) error {
	query := `UPDATE esp32 SET idUser = ? WHERE idESP32 = ?`
	
	_, err := r.db.ExecContext(ctx, query, userID, esp32ID)
	return err
}

// UnassignFromUser desasigna un ESP32 de cualquier usuario
func (r *MySQLESP32Repository) UnassignFromUser(ctx context.Context, esp32ID int) error {
	query := `UPDATE esp32 SET idUser = NULL WHERE idESP32 = ?`
	
	_, err := r.db.ExecContext(ctx, query, esp32ID)
	return err
}

// FindByNumeroSerie busca un ESP32 por su número de serie
func (r *MySQLESP32Repository) FindByNumeroSerie(ctx context.Context, numeroSerie string) (*entities.ESP32, error) {
	query := `SELECT idESP32, idKY_026, idMQ_2, idMQ_135, idDHT_22, numero_serie, idUser 
              FROM esp32 WHERE numero_serie = ?`
	
	var esp32 entities.ESP32
	var userID sql.NullInt64
	var idKY026 sql.NullInt64
	var idMQ2 sql.NullInt64
	var idMQ135 sql.NullInt64
	var idDHT22 sql.NullInt64
	
	err := r.db.QueryRowContext(ctx, query, numeroSerie).Scan(
		&esp32.ID,
		&idKY026,
		&idMQ2,
		&idMQ135,
		&idDHT22,
		&esp32.NumeroSerie,
		&userID,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No error, just no ESP32 found
		}
		return nil, err
	}
	
	// Convert nullable fields to int
	if idKY026.Valid {
		esp32.IDKY026 = int(idKY026.Int64)
	}
	if idMQ2.Valid {
		esp32.IDMQ2 = int(idMQ2.Int64)
	}
	if idMQ135.Valid {
		esp32.IDMQ135 = int(idMQ135.Int64)
	}
	if idDHT22.Valid {
		esp32.IDDHT22 = int(idDHT22.Int64)
	}
	
	if userID.Valid {
		userIDInt := int(userID.Int64)
		esp32.UserID = &userIDInt
	}
	
	// Set current time as created_at since it's not in the database
	esp32.CreatedAt = time.Now()
	
	return &esp32, nil
}
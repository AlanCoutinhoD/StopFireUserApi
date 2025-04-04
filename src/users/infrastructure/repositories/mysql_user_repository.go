package repositories

import (
	"context"
	"database/sql"
	"errors"

	"hex_go/src/users/domain/entities"
	"hex_go/src/users/domain/repositories"
)

// MySQLUserRepository implementa UserRepository usando MySQL
type MySQLUserRepository struct {
	db *sql.DB
}

// NewMySQLUserRepository crea una nueva instancia de MySQLUserRepository
func NewMySQLUserRepository(db *sql.DB) repositories.UserRepository {
	return &MySQLUserRepository{
		db: db,
	}
}

// Create inserta un nuevo usuario en la base de datos
func (r *MySQLUserRepository) Create(ctx context.Context, user *entities.User) (*entities.User, error) {
	query := `INSERT INTO users (username, password, email) VALUES (?, ?, ?)`
	
	result, err := r.db.ExecContext(ctx, query, user.Username, user.Password, user.Email)
	if err != nil {
		return nil, err
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	
	user.ID = int(id)
	
	return user, nil
}

// FindByID busca un usuario por su ID
func (r *MySQLUserRepository) FindByID(ctx context.Context, id int) (*entities.User, error) {
	query := `SELECT id, username, password, email, created_at FROM users WHERE id = ?`
	
	var user entities.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	
	return &user, nil
}

// FindByUsername busca un usuario por su nombre de usuario
func (r *MySQLUserRepository) FindByUsername(ctx context.Context, username string) (*entities.User, error) {
	query := `SELECT id, username, password, email, created_at FROM users WHERE username = ?`
	
	var user entities.User
	err := r.db.QueryRowContext(ctx, query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No error, just no user found
		}
		return nil, err
	}
	
	return &user, nil
}

// FindByEmail busca un usuario por su email
func (r *MySQLUserRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	query := `SELECT id, username, password, email, created_at FROM users WHERE email = ?`
	
	var user entities.User
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No error, just no user found
		}
		return nil, err
	}
	
	return &user, nil
}

// Update actualiza un usuario existente
func (r *MySQLUserRepository) Update(ctx context.Context, user *entities.User) error {
	query := `UPDATE users SET username = ?, password = ?, email = ? WHERE id = ?`
	
	_, err := r.db.ExecContext(ctx, query, user.Username, user.Password, user.Email, user.ID)
	return err
}

// Delete elimina un usuario por su ID
func (r *MySQLUserRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = ?`
	
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
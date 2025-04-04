package services

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"hex_go/src/users/domain/entities"
	"hex_go/src/users/domain/repositories"
)

// LoginUserUseCase implementa el caso de uso para iniciar sesión
type LoginUserUseCase struct {
	userRepository repositories.UserRepository
}

// NewLoginUserUseCase crea una nueva instancia de LoginUserUseCase
func NewLoginUserUseCase(userRepo repositories.UserRepository) *LoginUserUseCase {
	return &LoginUserUseCase{
		userRepository: userRepo,
	}
}

// LoginResponse contiene el token JWT y la información del usuario
type LoginResponse struct {
	Token string        `json:"token"`
	User  *entities.User `json:"user"`
}

// Execute ejecuta el caso de uso
func (uc *LoginUserUseCase) Execute(ctx context.Context, email, password string) (*LoginResponse, error) {
	// Buscar usuario por email
	user, err := uc.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	
	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	// Verificar contraseña
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Generar token JWT
	token, err := generateJWT(user)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Token: token,
		User:  user,
	}, nil
}

// generateJWT genera un token JWT para el usuario
func generateJWT(user *entities.User) (string, error) {
	// Obtener clave secreta de las variables de entorno
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "your_jwt_secret_key" // Valor por defecto
	}

	// Crear claims
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token válido por 24 horas
	}

	// Crear token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firmar token
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hex_go/src/users/application/services"
)

// UserController maneja las solicitudes HTTP para usuarios
type UserController struct {
	createUserUseCase *services.CreateUserUseCase
	loginUserUseCase  *services.LoginUserUseCase
}

// NewUserController crea una nueva instancia de UserController
func NewUserController(createUserUseCase *services.CreateUserUseCase, loginUserUseCase *services.LoginUserUseCase) *UserController {
	return &UserController{
		createUserUseCase: createUserUseCase,
		loginUserUseCase:  loginUserUseCase,
	}
}

// CreateUserRequest representa la estructura de la solicitud para crear un usuario
type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

// LoginRequest representa la estructura de la solicitud para iniciar sesión
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Register maneja la solicitud HTTP para registrar un nuevo usuario
func (c *UserController) Register(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.createUserUseCase.Execute(ctx, req.Username, req.Password, req.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

// Login maneja la solicitud HTTP para iniciar sesión
func (c *UserController) Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.loginUserUseCase.Execute(ctx, req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// SetupRoutes configura las rutas para el controlador de usuarios
func (c *UserController) SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("/register", c.Register)
			users.POST("/login", c.Login)
			users.GET("", func(ctx *gin.Context) {
				ctx.JSON(200, gin.H{
					"message": "Users API is running",
				})
			})
		}
	}
}
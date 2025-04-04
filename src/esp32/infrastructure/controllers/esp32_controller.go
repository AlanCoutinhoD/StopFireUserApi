package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"hex_go/src/esp32/application/services"
	"hex_go/src/esp32/domain/repositories"
)

// ESP32Controller maneja las solicitudes HTTP para ESP32
type ESP32Controller struct {
	assignESP32UseCase   *services.AssignESP32UseCase
	unassignESP32UseCase *services.UnassignESP32UseCase
	getUserESP32sUseCase *services.GetUserESP32sUseCase
	esp32Repository      repositories.ESP32Repository
}

// NewESP32Controller crea una nueva instancia de ESP32Controller
func NewESP32Controller(
	assignESP32UseCase *services.AssignESP32UseCase,
	unassignESP32UseCase *services.UnassignESP32UseCase,
	getUserESP32sUseCase *services.GetUserESP32sUseCase,
	esp32Repository repositories.ESP32Repository,
) *ESP32Controller {
	return &ESP32Controller{
		assignESP32UseCase:   assignESP32UseCase,
		unassignESP32UseCase: unassignESP32UseCase,
		getUserESP32sUseCase: getUserESP32sUseCase,
		esp32Repository:      esp32Repository,
	}
}

// AssignESP32Request representa la estructura de la solicitud para asignar un ESP32
type AssignESP32Request struct {
	NumeroSerie string `json:"numero_serie" binding:"required"`
}

// AssignESP32 maneja la solicitud HTTP para asignar un ESP32 a un usuario
func (c *ESP32Controller) AssignESP32(ctx *gin.Context) {
	// Obtener el ID del usuario del contexto (establecido por el middleware de autenticación)
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	var req AssignESP32Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Buscar el ESP32 por número de serie
	esp32, err := c.esp32Repository.FindByNumeroSerie(ctx, req.NumeroSerie)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if esp32 == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "ESP32 not found"})
		return
	}

	// Asignar el ESP32 al usuario
	err = c.assignESP32UseCase.Execute(ctx, esp32.ID, userID.(int))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "ESP32 assigned successfully"})
}

// UnassignESP32 maneja la solicitud HTTP para desasignar un ESP32 de un usuario
func (c *ESP32Controller) UnassignESP32(ctx *gin.Context) {
	// Obtener el ID del ESP32 de la URL
	esp32IDStr := ctx.Param("id")
	esp32ID, err := strconv.Atoi(esp32IDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ESP32 ID"})
		return
	}

	err = c.unassignESP32UseCase.Execute(ctx, esp32ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "ESP32 unassigned successfully"})
}

// GetUserESP32s maneja la solicitud HTTP para obtener todos los ESP32 de un usuario
func (c *ESP32Controller) GetUserESP32s(ctx *gin.Context) {
	// Obtener el ID del usuario del contexto (establecido por el middleware de autenticación)
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	esp32s, err := c.getUserESP32sUseCase.Execute(ctx, userID.(int))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, esp32s)
}

// SetupRoutes configura las rutas para el controlador de ESP32
func (c *ESP32Controller) SetupRoutes(router *gin.Engine, authMiddleware gin.HandlerFunc) {
	api := router.Group("/api")
	{
		esp32s := api.Group("/esp32s")
		{
			// Rutas protegidas (requieren autenticación)
			protected := esp32s.Group("")
			protected.Use(authMiddleware)
			{
				protected.POST("/assign", c.AssignESP32)
				protected.DELETE("/:id/unassign", c.UnassignESP32)
				protected.GET("/user", c.GetUserESP32s)
			}
		}
	}
}
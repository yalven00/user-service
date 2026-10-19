package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"user-service/internal/auth"
	"user-service/internal/config"
	"user-service/internal/user"
)

func New(db *gorm.DB, cfg *config.Config) *gin.Engine {
	r := gin.Default()

	// --- Health-check (публичный) ---
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// --- Сборка зависимостей ---
	repo := user.NewRepository(db)
	svc := user.NewService(repo, cfg.JWTSecret, cfg.JWTAccessExpiry)
	handler := user.NewHandler(svc)

	api := r.Group("/api/v1")

	// --- Публичные маршруты ---
	api.POST("/users", handler.Create)         // регистрация
	api.POST("/auth/login", handler.Login)     // логин → JWT

	// --- Защищённые маршруты (требуют Bearer-токен) ---
	protected := api.Group("")
	protected.Use(auth.JWTMiddleware(cfg.JWTSecret))
	{
		protected.GET("/auth/me", handler.Me)
		protected.GET("/users", handler.List)
		protected.GET("/users/:id", handler.GetByID)
		protected.PUT("/users/:id", handler.Update)
		protected.DELETE("/users/:id", handler.Delete)
	}

	return r
}
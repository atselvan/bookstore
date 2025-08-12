package app

import (
	"log/slog"

	"github.com/atselvan/bkst/domain/ports"
	"github.com/atselvan/bkst/repositories/user"
	"github.com/atselvan/bkst/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type application struct {
	usersSrv ports.UserService
}

func Start() error {
	// Initialize the user repository
	userRepo := user.NewUserRepository()

	// Initialize the user service with the user repository
	usersSrv := services.NewUserService(userRepo)

	// Create the application instance
	app := &application{
		usersSrv: usersSrv,
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/api/users/:id", app.handleGetUserByID)
	slog.Info("Server started on port 8080")
	if err := r.Run(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		return err
	}
	return nil
}

func (app *application) handleGetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(500, gin.H{"error": "internal error"})
		return
	}
	user, err := app.usersSrv.FindByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "internal error"})
		return
	}
	c.JSON(200, user)
}

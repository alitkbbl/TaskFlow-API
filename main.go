package main

import (
	"TaskFlowAPI/config"
	"TaskFlowAPI/handlers"
	"TaskFlowAPI/middlewares"
	"TaskFlowAPI/repository"
	"TaskFlowAPI/services"
	"TaskFlowAPI/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// DB connection
	db, err := repository.NewDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Repos
	userRepo := repository.NewUserRepository(db)
	taskRepo := repository.NewTaskRepository(db)

	// Utils
	jwtUtil := utils.NewJWTUtil(cfg.JWTSecret, cfg.JWTExpMinutes)

	// Services
	authService := services.NewAuthService(userRepo, jwtUtil)
	taskService := services.NewTaskService(taskRepo)

	// Handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	// Router
	r := gin.Default()
	r.Use(middlewares.Logging())
	r.Use(middlewares.CORSMiddleware())

	api := r.Group("/api/v1")
	{
		// Auth routes
		api.POST("/users/register", authHandler.Register)
		api.POST("/users/login", authHandler.Login)

		// Protected
		api.GET("/users/profile", middlewares.AuthMiddleware(jwtUtil), userHandler.Profile)
		api.PUT("/users/profile", middlewares.AuthMiddleware(jwtUtil), userHandler.UpdateProfile)

		// Tasks
		api.GET("/tasks", taskHandler.ListTasks)
		api.GET("/tasks/:id", taskHandler.GetTask)

		api.POST("/tasks", middlewares.AuthMiddleware(jwtUtil), taskHandler.CreateTask)
		api.PUT("/tasks/:id", middlewares.AuthMiddleware(jwtUtil), taskHandler.UpdateTask)
		api.DELETE("/tasks/:id", middlewares.AuthMiddleware(jwtUtil), taskHandler.DeleteTask)
		api.PUT("/tasks/:id/status", middlewares.AuthMiddleware(jwtUtil), taskHandler.UpdateStatus)
	}

	log.Printf("Server running on port %s", cfg.Port)
	r.Run(":" + cfg.Port)
}

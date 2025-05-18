package main

import (
	"diploma/internal/config"
	"diploma/internal/order"
	"diploma/internal/user"
	"diploma/pkg/storage/postgres"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg := config.Load()

	storage, err := postgres.InitStorage(cfg.DatabaseDSN)
	if err != nil {
		log.Fatal(err)
	}

	repo := storage
	userHandler := user.NewHandler(repo, cfg.JWTSecret)
	orderHandler := order.NewHandler(repo)

	r := gin.Default()
	r.POST("/api/user/register", userHandler.Register)
	r.POST("/api/user/login", userHandler.Login)
	r.POST("/api/user/orders", user.AuthMiddleware(cfg.JWTSecret), orderHandler.CreateOrder)
	r.GET("/api/user/orders", user.AuthMiddleware(cfg.JWTSecret), orderHandler.GetOrders)

	r.Run()
}

package main

import (
    "log"

    "github.com/gin-gonic/gin"

    "github.com/hugaojanuario/crud-golang/config"
    "github.com/hugaojanuario/crud-golang/internal/database/postgre"
    "github.com/hugaojanuario/crud-golang/internal/user"
)

func main() {
	cfg := config.Load()

	db, err := postgre.NewConnection(postgre.Config{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		DBName:   cfg.DBName,
		SSLMode:  cfg.DBSSLMode,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer db.Close()

	repo := user.NewRepository(db)
	service := user.NewService(repo)
	handler := user.NewHandler(service)

	r := gin.Default()
	user.RegisterRoutes(r, handler)

	log.Printf("servidor rodando na porta %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf(err.Error())
	}
}

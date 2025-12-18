package main

import (
	"net/http"

	"go-rest-api/internal/auth"
	"go-rest-api/internal/config"
	"go-rest-api/internal/db"
	"go-rest-api/internal/logger"
	"go-rest-api/internal/user"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

func main() {
	logger.Init()

	cfg := config.Config{
		DBHost:     "localhost",
		DBPort:     5433,
		DBUser:     "postgres",
		DBPassword: "postgres",
		DBName:     "testdb",
		SSLMode:    "disable",
		JWTKey:     "SuperSecretKey123456",
	}

	database, err := db.NewPostgres(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("DB connection failed")
	}
	defer database.Close()

	userRepo := user.NewPgRepository(database)
	userValidator := user.NewValidator()
	userService := user.NewService(userRepo, userValidator)
	userHandler := user.NewHandler(userService)

	r := chi.NewRouter()

	r.Post("/login", auth.Login)

	r.Route("/api/users", func(rt chi.Router) {
		rt.Use(auth.Middleware) // защищенный доступ
		rt.Mount("/", userHandler.Routes())
	})

	log.Info().Msg("Server started on :8080")
	http.ListenAndServe(":8080", r)
}

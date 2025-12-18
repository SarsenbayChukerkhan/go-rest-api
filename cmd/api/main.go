package main

import (
	"net/http"
	"os"

	"go-rest-api/internal/auth"
	"go-rest-api/internal/db"
	"go-rest-api/internal/logger"
	"go-rest-api/internal/user"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

func main() {
	// инициализация логирования
	logger.Init()

	// подключение к PostgreSQL через DATABASE_URL (Render)
	database, err := db.NewPostgresFromEnv()
	if err != nil {
		log.Fatal().Err(err).Msg("DB connection failed")
	}
	defer database.Close()

	// Dependency Injection
	userRepo := user.NewPgRepository(database)
	userValidator := user.NewValidator()
	userService := user.NewService(userRepo, userValidator)
	userHandler := user.NewHandler(userService)

	// роутер
	r := chi.NewRouter()

	// login (JWT)
	r.Post("/login", auth.Login)

	// защищённые маршруты
	r.Route("/api/users", func(rt chi.Router) {
		rt.Use(auth.Middleware)
		rt.Mount("/", userHandler.Routes())
	})

	// PORT от Render
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Info().Msg("Server started on :" + port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal().Err(err).Msg("Server failed")
	}
}

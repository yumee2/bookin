package main

import (
	"log"
	"yumee/booking/internal/booking"
	"yumee/booking/internal/config"
	"yumee/booking/internal/httpapi"
	"yumee/booking/internal/postgres"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	repo, err := postgres.NewRepository(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("connect to postgres: %v", err)
	}

	service := booking.NewBookingService(repo)
	router := httpapi.NewRouter(service)

	if err := router.Run(cfg.HTTPAddr); err != nil {
		log.Fatal(err)
	}
}

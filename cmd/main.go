package main

import (
	"log"
	"yumee/booking/internal/booking"
	"yumee/booking/internal/httpapi"
	"yumee/booking/internal/postgres"
)

// move it to config later
const dsn = "host=localhost user=bookin password=2314 dbname=bookin port=5432 sslmode=disable"

func main() {
	repo, err := postgres.NewRepository(dsn)
	if err != nil {
		log.Fatalf("connect to postgres: %v", err)
	}

	service := booking.NewBookingService(repo)
	router := httpapi.NewRouter(service)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

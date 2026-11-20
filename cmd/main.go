package main

import (
	"log"
	"yumee/booking/internal/httpapi"
	"yumee/booking/internal/postgres"
)

// move it to config later
const dsn = "host=localhost user=bookin password=2314 dbname=bookin port=5432 sslmode=disable"

func main() {
	_, err := postgres.NewRepository(dsn)
	if err != nil {
		log.Fatalf("Error trying to connect to psql: ", err)
	}

	ginEngine := httpapi.NewRouter()
	if err := ginEngine.Run(); err != nil {
		log.Fatal(err)
	}
}

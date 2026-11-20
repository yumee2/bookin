package main

import (
	"yumee/booking/internal/httpapi"
)

func main() {
	genEngine := httpapi.NewRouter()
	genEngine.Run()
}

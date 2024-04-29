package main

import (
	"raha-beach-code-challenge/internal/itinerary"
	"raha-beach-code-challenge/pkg/http"
)

func main() {
	e := http.NewServer()
	itinerary.RegisterHandlers(e)
	e.Logger.Fatal(e.Start(":1323"))
}

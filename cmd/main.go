package main

// Import dependencies that are part of this project.
import (
	"raha-beach-code-challenge/internal/itinerary" // Imports functions and types related to itinerary management.
	"raha-beach-code-challenge/pkg/http"           // Imports the HTTP server setup utilities.
)

// The main function is the entry point of the application.
func main() {

	// Create a new HTTP server using predefined settings from the http package.
	e := http.NewServer()

	// Register API endpoints related to itineraries, linking them to the server.
	itinerary.RegisterHandlers(e)

	// Start listening for HTTP requests on port 1323.
	// If there's an issue starting the server, it will log the error and exit the program.
	e.Logger.Fatal(e.Start(":1323"))
}

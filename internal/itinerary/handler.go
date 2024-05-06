package itinerary

// Import necessary libraries and packages.
import (
	"net/http" // Standard library for HTTP server functionality.

	"github.com/labstack/echo/v4" // Echo framework for handling HTTP requests.
)

// RegisterHandlers sets up the routes and links them to handler functions.
func RegisterHandlers(e *echo.Echo) {
	// Register the POST method for the "/itinerary" route and link it to the handleItinerary function.
	e.POST("/itineraries", handleItinerary)
	e.GET("/airports", listAirportsHandler)
	e.GET("/airports/count", countAirportsHandler)
}

func listAirportsHandler(c echo.Context) error {
	service := NewService()
	airports := service.ListAirports()
	return c.JSON(http.StatusOK, airports)
}

func countAirportsHandler(c echo.Context) error {
	service := NewService()
	count := service.CountAirports()
	return c.JSON(http.StatusOK, map[string]int{"count": count})
}

// handleItinerary processes POST requests to compute an itinerary from a series of ticket pairs.
func handleItinerary(c echo.Context) error {
	var tickets []TicketPair // Declare a slice to store the incoming ticket pairs.

	// Bind the incoming JSON payload to the tickets slice. If there's an error, return a bad request error.
	if err := c.Bind(&tickets); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	// Create a new instance of the service that processes itineraries.
	service := NewService()

	// Call the ReconstructItinerary method to process the tickets and reconstruct the itinerary.
	itinerary, err := service.ReconstructItinerary(tickets)

	// If there's an error during the process, return an internal server error.
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not reconstruct itinerary")
	}

	// If successful, return the itinerary as a JSON response with HTTP status 200 OK.
	return c.JSON(http.StatusOK, itinerary)
}

package itinerary

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterHandlers(e *echo.Echo) {
	e.POST("/itinerary", handleItinerary)
}

func handleItinerary(c echo.Context) error {
	var tickets []TicketPair
	if err := c.Bind(&tickets); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	service := NewService()
	itinerary, err := service.ReconstructItinerary(tickets)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not reconstruct itinerary")
	}

	return c.JSON(http.StatusOK, itinerary)
}

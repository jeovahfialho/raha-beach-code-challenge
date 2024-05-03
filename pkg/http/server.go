package http

// Import necessary packages from external libraries.
import (
	"github.com/labstack/echo/v4"            // Import the Echo framework for handling HTTP requests.
	"github.com/labstack/echo/v4/middleware" // Import Echo's built-in middleware for enhanced functionality.
)

// NewServer function initializes and returns a new Echo server. This server will be used
// throughout the application to handle HTTP requests.
func NewServer() *echo.Echo {
	e := echo.New() // Create a new instance of an Echo server.

	// Middleware
	e.Use(middleware.Logger())  // Attach Logger middleware which logs all incoming requests.
	e.Use(middleware.Recover()) // Attach Recover middleware which helps to recover from panics anywhere in the stack chain and handles the HTTP response accordingly.

	return e // Return the initialized Echo server instance.
}

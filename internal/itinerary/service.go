package itinerary

// Import necessary standard libraries.
import (
	"errors" // Used for creating error objects.
	"fmt"    // Used for formatting strings.
)

// Service struct defines a service that offers itinerary reconstruction functionality.
type Service struct{}

// NewService is a constructor function that initializes and returns a new instance of Service.
func NewService() *Service {
	return &Service{} // Return a new instance of Service.
}

// ReconstructItinerary takes an array of ticket pairs and attempts to reconstruct the complete travel itinerary.
func (s *Service) ReconstructItinerary(tickets []TicketPair) ([]string, error) {
	// Check if no tickets are provided. If so, return an error immediately.
	if len(tickets) == 0 {
		return nil, errors.New("no tickets provided")
	}

	// Initialize a map to track flight routes from source to destination.
	flightMap := make(map[string]string)
	// Initialize a map to track all destinations.
	destinations := make(map[string]bool)

	// Iterate through each ticket pair in the tickets array.
	for _, ticket := range tickets {
		// Check if the source airport already exists in the flightMap, which indicates a problem in the ticket chain.
		if _, exists := flightMap[ticket[0]]; exists {
			return nil, fmt.Errorf("invalid ticket chain: multiple entries for %s", ticket[0])
		}
		// Map the source airport to the destination airport.
		flightMap[ticket[0]] = ticket[1]
		// Mark the destination airport in the destinations map.
		destinations[ticket[1]] = true
	}

	// Attempt to find the starting airport of the itinerary.
	start := ""
	for src := range flightMap {
		// Check if the source is not listed as a destination.
		if !destinations[src] {
			// If a start point has already been found, this means there are multiple start points, which is an error.
			if start != "" {
				return nil, errors.New("multiple possible start points")
			}
			start = src // Set the starting point.
		}
	}
	// If no start is found, return an error because the itinerary cannot begin.
	if start == "" {
		return nil, errors.New("cannot find the start of the itinerary")
	}

	// Begin reconstructing the itinerary from the starting point.
	itinerary := []string{start}
	current := start
	usedTickets := make(map[string]bool)

	// Continue building the itinerary until no further destinations are found.
	for {
		next, exists := flightMap[current]
		if !exists || usedTickets[current] {
			break // Exit the loop if there is no next destination or the current ticket has already been used.
		}
		itinerary = append(itinerary, next) // Add the next destination to the itinerary.
		usedTickets[current] = true         // Mark the current ticket as used.
		current = next                      // Move to the next destination.
	}

	// After constructing the itinerary, ensure that it includes all tickets plus the starting point.
	if len(itinerary) != len(tickets)+1 {
		return nil, errors.New("itinerary cannot be completed from the given tickets")
	}

	return itinerary, nil // Return the successfully reconstructed itinerary.
}

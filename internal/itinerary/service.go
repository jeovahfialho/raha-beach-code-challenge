package itinerary

import (
	"errors"
	"fmt"
)

// Service provides itinerary reconstruction functionality.
type Service struct{}

// NewService creates a new Service.
func NewService() *Service {
	return &Service{}
}

// ReconstructItinerary reconstructs the complete travel itinerary from an array of ticket pairs.
func (s *Service) ReconstructItinerary(tickets []TicketPair) ([]string, error) {
	if len(tickets) == 0 {
		return nil, errors.New("no tickets provided")
	}

	flightMap := make(map[string]string)
	destinations := make(map[string]bool)

	for _, ticket := range tickets {
		if _, exists := flightMap[ticket[0]]; exists {
			return nil, fmt.Errorf("invalid ticket chain: multiple entries for %s", ticket[0])
		}
		flightMap[ticket[0]] = ticket[1]
		destinations[ticket[1]] = true
	}

	// Find the starting point
	start := ""
	for src := range flightMap {
		if !destinations[src] {
			if start != "" {
				return nil, errors.New("multiple possible start points")
			}
			start = src
		}
	}
	if start == "" {
		return nil, errors.New("cannot find the start of the itinerary")
	}

	// Reconstruct the itinerary
	itinerary := []string{start}
	current := start
	usedTickets := make(map[string]bool)

	for {
		next, exists := flightMap[current]
		if !exists || usedTickets[current] {
			break
		}
		itinerary = append(itinerary, next)
		usedTickets[current] = true
		current = next
	}

	if len(itinerary) != len(tickets)+1 {
		return nil, errors.New("itinerary cannot be completed from the given tickets")
	}

	return itinerary, nil
}

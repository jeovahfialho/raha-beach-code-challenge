package itinerary

import (
	"errors"
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

	// Find the starting point
	start, err := findStart(tickets)
	if err != nil {
		return nil, err
	}

	// Reconstruct the itinerary
	itinerary := []string{start}
	for {
		if len(itinerary)-1 == len(tickets) {
			break
		}
		next, err := findNextDestination(itinerary[len(itinerary)-1], tickets)
		if err != nil {
			return nil, err
		}
		itinerary = append(itinerary, next)
	}

	return itinerary, nil
}

// findStart finds the starting airport code.
func findStart(tickets []TicketPair) (string, error) {
	destinationCounts := make(map[string]int)
	for _, pair := range tickets {
		destinationCounts[pair[1]]++
	}

	for _, pair := range tickets {
		if destinationCounts[pair[0]] == 0 {
			return pair[0], nil
		}
	}
	return "", errors.New("cannot find the start of the itinerary")
}

// findNextDestination finds the next destination given a source.
func findNextDestination(source string, tickets []TicketPair) (string, error) {
	for _, pair := range tickets {
		if pair[0] == source {
			return pair[1], nil
		}
	}
	return "", errors.New("next destination not found")
}

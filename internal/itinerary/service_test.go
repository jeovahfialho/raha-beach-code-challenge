package itinerary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReconstructItinerary(t *testing.T) {
	tests := []struct {
		name           string
		tickets        []TicketPair
		expected       []string
		expectingError bool
	}{
		{
			name: "Simple itinerary",
			tickets: []TicketPair{
				{"JFK", "LAX"},
				{"LAX", "DXB"},
				{"DXB", "SFO"},
				{"SFO", "SJC"},
			},
			expected: []string{"JFK", "LAX", "DXB", "SFO", "SJC"},
		},
		{
			name: "Disjoint itinerary",
			tickets: []TicketPair{
				{"ATL", "EWR"},
				{"SFO", "ATL"},
				{"GSO", "SFO"},
				{"EWR", "GSO"},
			},
			expected:       []string{"EWR", "GSO", "SFO", "ATL", "EWR"},
			expectingError: true, // This should error out because it is disjoint and creates a loop.
		},
		{
			name:           "No tickets",
			tickets:        []TicketPair{},
			expectingError: true,
		},
		{
			name: "Invalid ticket chain",
			tickets: []TicketPair{
				{"ORD", "DEN"},
				{"DEN", "SFO"},
				// Missing link from SFO
			},
			expectingError: true,
		},
		{
			name: "Single ticket",
			tickets: []TicketPair{
				{"MIA", "LHR"},
			},
			expected: []string{"MIA", "LHR"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			service := NewService()
			actual, err := service.ReconstructItinerary(test.tickets)
			if test.expectingError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, actual)
			}
		})
	}
}

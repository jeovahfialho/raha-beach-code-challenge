package itinerary

// Import necessary libraries.
import (
	"testing" // Import Go's testing framework to enable writing test functions.

	"github.com/stretchr/testify/assert" // Import the testify library to provide more assert functions for testing.
)

// TestReconstructItinerary tests the functionality of the ReconstructItinerary method.
func TestReconstructItinerary(t *testing.T) {
	// Define a list of test cases.
	tests := []struct {
		name           string       // Name of the test case.
		tickets        []TicketPair // Input tickets for the test.
		expected       []string     // Expected result from the test.
		expectingError bool         // Whether an error is expected.
	}{
		// Test case: Simple itinerary without loops or breaks.
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
		// Test case: Itinerary that ends up forming a loop, which should result in an error.
		{
			name: "Disjoint itinerary",
			tickets: []TicketPair{
				{"ATL", "EWR"},
				{"SFO", "ATL"},
				{"GSO", "SFO"},
				{"EWR", "GSO"},
			},
			expected:       []string{"EWR", "GSO", "SFO", "ATL", "EWR"},
			expectingError: true, // This case is expected to produce an error due to the loop.
		},
		// Test case: No tickets provided, should typically error out.
		{
			name:           "No tickets",
			tickets:        []TicketPair{},
			expectingError: true,
		},
		// Test case: A single ticket, straightforward scenario.
		{
			name: "Single ticket",
			tickets: []TicketPair{
				{"MIA", "LHR"},
			},
			expected: []string{"MIA", "LHR"},
		},
	}

	// Iterate through each test case.
	for _, test := range tests {
		// Subtest for each scenario.
		t.Run(test.name, func(t *testing.T) {
			// Instantiate the service that will reconstruct the itinerary.
			service := NewService()
			// Attempt to reconstruct the itinerary based on the test's ticket pairs.
			actual, err := service.ReconstructItinerary(test.tickets)
			// Check if an error was expected.
			if test.expectingError {
				// If an error was expected, assert that an error occurred.
				assert.Error(t, err)
			} else {
				// If no error was expected, assert that no error occurred and that the results match expectations.
				assert.NoError(t, err)
				assert.Equal(t, test.expected, actual)
			}
		})
	}
}

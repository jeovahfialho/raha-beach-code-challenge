package itinerary

// TicketPair is defined as a type alias for an array of two strings. This type is used to represent a pair of locations:
// the departure and the arrival airports. Each TicketPair consists of two elements:
// [0] - The departure airport code (e.g., "JFK")
// [1] - The arrival airport code (e.g., "LAX")
// Using a type alias here improves code readability by clearly indicating the intended use of these string arrays
// throughout the itinerary package.
type TicketPair [2]string

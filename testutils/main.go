// Define the package main, indicating that this file can be compiled into an executable program.
package main

// Import necessary standard libraries.
import (
	"bytes"         // Used for interacting with byte slices.
	"encoding/json" // For encoding and decoding JSON.
	"flag"          // For parsing command-line flags.
	"fmt"           // For formatted I/O operations.
	"io/ioutil"     // For reading and writing files.
	"math/rand"     // For generating random numbers.
	"net/http"      // For making HTTP requests.
	"os"            // For interacting with the operating system.
	"time"          // For time-related operations.
)

// generateTestData generates test data with a specified number of ticket pairs.
func generateTestData(numTickets int) []byte {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator.

	// Initialize a list of airport codes.
	airports := []string{
		"JFK", "LAX", "DXB", "SFO", "SJC", "ORD", "ATL", "PEK", "CDG", "AMS",
		"FRA", "DEN", "ICN", "MAD", "LAS", "BCN", "MUC", "SEA", "FCO", "BKK",
		"MCO", "EWR", "IST", "YYZ", "SYD", "IAH", "CAN", "DEL", "MIA", "MEX",
		"PHX", "LGA", "MNL", "CGK", "SZX", "MSP", "PHL", "BOS", "CTU", "PVG",
		"IAD", "ABC", "MEL", "DUB", "CLT", "ZRH", "MAN", "TXL", "KUL", "BRU",
		"OSL", "ARN", "ATH", "VIE", "CPH", "HAM", "NRT", "HEL", "PMI", "SAW",
		"EDI", "CAI", "KMG", "CGO", "CSX", "CJU", "XMN", "WUH", "XIY", "TAO",
		"TPE", "KHH", "CNX", "BKI", "COK", "HKT", "USM", "DPS", "SUB", "OKA",
		"KIX", "HND", "BOM", "MAA", "CCU", "HYD", "AMD", "BLR", "GOI", "TRV",
	}

	// Generate additional synthetic airport codes to expand the dataset.
	for i := 0; i < 900; i++ {
		airports = append(airports, fmt.Sprintf("A%d", i+1))
	}

	var tickets [][]string

	// Generate ticket pairs ensuring they are sequentially connected without forming a loop.
	for i := 0; i < numTickets-1; i++ {
		src := airports[i%len(airports)]
		dst := airports[(i+1)%len(airports)]
		tickets = append(tickets, []string{src, dst})
	}

	// Marshal the slice of ticket pairs into JSON format.
	ticketsJSON, _ := json.Marshal(tickets)
	return ticketsJSON
}

// runLoadTest performs a load test by sending JSON data to a specified endpoint.
func runLoadTest(endpointURL string, jsonData []byte) {
	startTime := time.Now()                                                            // Record the start time of the test.
	resp, err := http.Post(endpointURL, "application/json", bytes.NewBuffer(jsonData)) // Make a POST request with JSON data.
	if err != nil {
		panic(err) // Handle errors by panicking.
	}
	defer resp.Body.Close() // Ensure the response body is closed after the function exits.

	body, _ := ioutil.ReadAll(resp.Body) // Read the entire response body.
	elapsedTime := time.Since(startTime) // Calculate the elapsed time since the start of the test.
	fmt.Println("Response time:", elapsedTime)
	fmt.Println("Server response:", string(body))
}

// main function interprets command-line arguments and runs specified utilities.
func main() {
	var genData bool
	var runTest bool
	var numTickets int
	var endpointURL string

	flag.BoolVar(&genData, "generate", false, "Generate test data")
	flag.BoolVar(&runTest, "loadtest", false, "Run load test")
	flag.IntVar(&numTickets, "tickets", 100, "Number of tickets to generate")
	flag.StringVar(&endpointURL, "url", "http://localhost:1323/itinerary", "Endpoint URL for load testing")
	flag.Parse() // Parse command-line flags.

	if genData {
		data := generateTestData(numTickets) // Generate data if requested.
		os.Stdout.Write(data)                // Output the generated data to standard output.
	}

	if runTest {
		jsonData, err := ioutil.ReadFile("./generated_data.json") // Read the generated JSON data from a file.
		if err != nil {
			fmt.Println("Error reading JSON data:", err)
			return
		}
		runLoadTest(endpointURL, jsonData) // Run a load test using the read data.
	}
}

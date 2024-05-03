
# Itinerary Reconstructor

## Overview
This Go web application uses the Echo framework to provide an HTTP endpoint that reconstructs a complete flight itinerary from a series of flight tickets represented as JSON pairs.

## Installation

### Option 1: Running Locally

To run this application, you need Go installed on your machine.

1. Install Go from [the official website](https://golang.org/dl/).
2. Clone the repository to your local machine.
3. Navigate to the root directory of the project.
4. Run `go mod tidy` to install the dependencies.
5. Start the server with `go run ./cmd/main.go`.

### Option 2: Running with Docker

1. Ensure you have Docker installed on your system.
2. Clone the repository to your local machine.
3. Navigate to the root directory of the project.
4. Build the Docker image using the Dockerfile:

```bash
docker build -t itinerary-reconstructor .
```

## Project Structure
The project follows a clean architecture that separates concerns, making the code more maintainable and scalable.

- `/cmd` contains the application entry point.
- `/internal` contains business logic.
- `/pkg` contains reusable packages across applications.

## SOLID Principles
The codebase adheres to SOLID principles:

- **Single Responsibility**: `service.go` handles the business logic of itinerary reconstruction, while `handler.go` manages HTTP requests and responses.
- **Open/Closed**: The application is open for extension but closed for modification. New functionality can be added without changing existing code.
- **Liskov Substitution**: Interfaces are designed to be replaceable.
- **Interface Segregation**: The application uses narrow, specific interfaces.
- **Dependency Inversion**: The application depends on abstractions, not on concretions.

## Code Quality
- **Clean Code**: Code is readable, simple, and concise.
- **Well-organized**: Files and packages are organized logically.
- **Properly Commented**: Crucial functions and complex logic are explained with comments.

## Performance
- **Efficient Algorithm**: The itinerary reconstruction is designed to be efficient and fast.
- **Minimalistic Framework**: Echo is a high-performance framework that adds minimal overhead.

## Testing
Comprehensive unit tests ensure the application meets all functional requirements.

- Run tests with `go test ./...`.

## Fulfilling Requirements and Criteria
- **Single Endpoint**: `internal/itinerary/handler.go` - `handleItinerary` function.
- **Input Format**: `internal/itinerary/types.go` - `TicketPair` type.
- **Functionality**: `internal/itinerary/service.go` - `ReconstructItinerary` method.
- **Expected Output**: `internal/itinerary/service.go` - `ReconstructItinerary` method.
- **Testing**: `internal/itinerary/service_test.go` - `TestReconstructItinerary` method.
- **Deliverables**: All source files, configuration files, and `README.md`.

### Acceptance Criteria Mapped
1. **Functionality**: Validated by unit tests in `service_test.go` and endpoint functionality in `handler.go`.
2. **Code Quality**: Demonstrated by the project's organized structure and clear code documentation.
3. **Performance**: Ensured by the Echo framework used in `server.go` and the algorithm in `service.go`.

## Running the Application
To start the application:

```
go run ./cmd/main.go
```

Then you can send a POST request to `http://localhost:1323/itinerary` with the JSON payload of ticket pairs.

## Example Request

```
curl -X POST http://localhost:1323/itinerary -H "Content-Type: application/json" -d '[["JFK","LAX"],["LAX","DXB"],["DXB","SFO"],["SFO","SJC"]]'
```

## Example Response

```
["JFK","LAX","DXB","SFO","SJC"]
```

# Performance Testing

## Generating Test Data

This utility allows for the generation of large-scale test data to evaluate the performance of the itinerary reconstruction service. The data consists of randomly generated flight ticket pairs.

### Expanded List of Airport Codes

The script includes an expanded list of fictitious airport codes beyond real-world major airport codes to simulate scenarios with a large dataset.

### Generating Tickets

The tickets are generated such that each ticket pair is connected end-to-end in a shuffled order to simulate a randomized set of flight connections without creating loops. This setup helps in simulating the real-world scenario where flight itineraries might need to be reconstructed from shuffled data.

To generate test data, navigate to the `testutils` directory and run the following command:

    go run ./main.go --generate --tickets 900 > ./generated_data.json

This command generates 900 ticket pairs and saves them in `generated_data.json`.

## Load Testing

The load testing script sends the generated ticket data as a JSON payload to the itinerary reconstruction endpoint and measures the response time and server's response to assess the application's performance under load.

### Running Load Test

After generating the data, you can perform the load test by running:

    go run ./main.go --loadtest --url http://localhost:1323/itinerary

This command reads the generated JSON data from `generated_data.json` and posts it to the specified URL, which is the itinerary endpoint. The script then outputs the response time and the server response, allowing you to assess how quickly and accurately the server processes the large input.

## What This Does

1. **Generate TestData**: The `generateTestData` function creates a large number of ticket pairs using an extensive list of airport codes. The function ensures that the list of tickets is shuffled to provide varied starting points and paths for the reconstruction algorithm, which mimics a real-world input scenario where data may not be sequentially ordered.

2. **Run Load Test**: The `runLoadTest` function performs a POST request with the generated ticket data to the server's itinerary endpoint. It measures how long the server takes to respond and prints out the response time and the actual JSON response from the server. This helps in identifying the performance characteristics of the application, such as response times and accuracy of the itinerary reconstruction.


Thank you for using the Itinerary Reconstructor.

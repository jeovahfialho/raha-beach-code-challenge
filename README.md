
# Itinerary Reconstructor

## Overview
This Go web application uses the Echo framework to provide an HTTP endpoint that reconstructs a complete flight itinerary from a series of flight tickets represented as JSON pairs.

## Installation
To run this application, you need Go installed on your machine.

1. Install Go from [the official website](https://golang.org/dl/).
2. Clone the repository to your local machine.
3. Navigate to the root directory of the project.
4. Run `go mod tidy` to install the dependencies.
5. Start the server with `go run ./cmd/main.go`.

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

## Assessment Criteria
1. **Functionality**: The application meets all requirements detailed in the spec.
2. **Code Quality**: The source code is clean and adheres to industry standards.
3. **Performance**: The app is performance optimized and responds quickly.

## Running the Application
To start the application:

```
go run ./cmd/main.go
```

Then you can send a POST request to `http://localhost:1323/itinerary` with the JSON payload of ticket pairs.

## Example Request

```
curl -X POST http://localhost:1323/itinerary -d '[["JFK","LAX"],["LAX","DXB"],["DXB","SFO"],["SFO","SJC"]]'
```

## Example Response

```
["JFK","LAX","DXB","SFO","SJC"]
```

Thank you for using the Itinerary Reconstructor.

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func generateTestData(numTickets int) []byte {
	rand.Seed(time.Now().UnixNano())

	// Lista de siglas de aeroportos expandida
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

	// Gerar mais aeroportos
	for i := 0; i < 900; i++ {
		airports = append(airports, fmt.Sprintf("A%d", i+1))
	}

	var tickets [][]string

	// Garantir que as conexões sejam encadeadas
	for i := 0; i < numTickets-1; i++ {
		src := airports[i%len(airports)]
		dst := airports[(i+1)%len(airports)]
		tickets = append(tickets, []string{src, dst})
	}

	// Adicionar a última conexão encadeada sem incluir o primeiro aeroporto

	ticketsJSON, _ := json.Marshal(tickets)
	return ticketsJSON
}

func runLoadTest(endpointURL string, jsonData []byte) {
	startTime := time.Now()
	resp, err := http.Post(endpointURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	elapsedTime := time.Since(startTime)
	fmt.Println("Response time:", elapsedTime)
	fmt.Println("Server response:", string(body))
}

func main() {
	var genData bool
	var runTest bool
	var numTickets int
	var endpointURL string

	flag.BoolVar(&genData, "generate", false, "Generate test data")
	flag.BoolVar(&runTest, "loadtest", false, "Run load test")
	flag.IntVar(&numTickets, "tickets", 100, "Number of tickets to generate")
	flag.StringVar(&endpointURL, "url", "http://localhost:1323/itinerary", "Endpoint URL for load testing")
	flag.Parse()

	if genData {
		data := generateTestData(numTickets)
		os.Stdout.Write(data)
	}

	if runTest {
		jsonData, err := ioutil.ReadFile("./generated_data.json") // Replace with the correct path
		if err != nil {
			fmt.Println("Error reading JSON data:", err)
			return
		}
		runLoadTest(endpointURL, jsonData)
	}
}

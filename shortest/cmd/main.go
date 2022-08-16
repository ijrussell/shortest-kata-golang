package main

import (
	"fmt"
	"os"
	"shortest"
)

func main() {
	filePath := "../resources/data.csv"
	start := "Cogburg"
	finish := "Leverstorm"
	// filePath := os.Args[1]
	// start := os.Args[2]
	// finish := os.Args[3]
	shortest, err := shortest.GetShortestRoute(filePath, start, finish)
	if err != nil {
		fmt.Printf("Error: %q", err)
		os.Exit(1)
	}
	fmt.Printf("%q, %d", append(shortest.Route, shortest.Location), shortest.TotalDistance)
}

// go run main.go
// go run main.go data.csv Cogburg Leverstorm

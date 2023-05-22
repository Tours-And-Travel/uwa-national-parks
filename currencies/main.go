package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
)

func main() {
	data := map[string]string{
		"FNR": "USD",
		"FR":  "USD",
		"EAC": "UGX",
	}

	// Convert data to JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal("Error converting data to JSON:", err)
	}

	// Write JSON data to a file
	jsonFile, err := os.Create("currencies.json")
	if err != nil {
		log.Fatal("Error creating JSON file:", err)
	}
	defer jsonFile.Close()

	_, err = jsonFile.Write(jsonData)
	if err != nil {
		log.Fatal("Error writing JSON data to file:", err)
	}

	log.Println("Conversion to JSON is complete. The output file is: currencies.json")

	// Create a new CSV file
	csvFile, err := os.Create("currencies.csv")
	if err != nil {
		log.Fatal("Unable to create file:", err)
	}
	defer csvFile.Close()

	// Create a CSV writer
	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Write CSV header
	header := []string{"Nationality", "Currency"}
	if err := writer.Write(header); err != nil {
		log.Fatal("Error writing CSV header:", err)
	}

	// Write CSV data
	for abbreviation, currency := range data {
		record := []string{abbreviation, currency}
		if err := writer.Write(record); err != nil {
			log.Fatal("Error writing CSV data:", err)
		}
	}

	log.Println("Conversion to CSV is complete. The output file is: currencies.csv")
}

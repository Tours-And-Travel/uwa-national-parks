package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
)

func main() {
	data := map[string]string{
		"Kidepo":               "571",
		"Bwindi":               "512",
		"Mgahinga":             "482",
		"Rwenzori":             "398",
		"Queen Elizabeth":      "389",
		"Semuliki":             "387",
		"Kibale":               "358",
		"Murchison":            "305",
		"Lake Mburo":           "253",
		"Mt.Elgon":             "229",
		"Ziwa Rhino Sanctuary": "164",
		"Pian Upe":             "402",
		"Toro semliki":         "325",
		"Katonga":              "274",
	}

	// Convert data to JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal("Error converting data to JSON:", err)
	}

	// Write JSON data to a file
	jsonFile, err := os.Create("distance.json")
	if err != nil {
		log.Fatal("Error creating JSON file:", err)
	}
	defer jsonFile.Close()

	_, err = jsonFile.Write(jsonData)
	if err != nil {
		log.Fatal("Error writing JSON data to file:", err)
	}

	log.Println("Conversion to JSON is complete. The output file is: distance.json")

	// Convert data to CSV
	csvFile, err := os.Create("distance.csv")
	if err != nil {
		log.Fatal("Error creating CSV file:", err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Write CSV header
	header := []string{"Park", "Distance (km)"}
	if err := writer.Write(header); err != nil {
		log.Fatal("Error writing CSV header:", err)
	}

	// Write CSV data
	for park, code := range data {
		record := []string{park, code}
		if err := writer.Write(record); err != nil {
			log.Fatal("Error writing CSV data:", err)
		}
	}

	log.Println("Conversion to CSV is complete. The output file is: distance.csv")
}

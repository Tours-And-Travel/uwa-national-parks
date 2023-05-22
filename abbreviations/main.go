package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
)

func main() {
	data := map[string]string{
		"FNR":   "Foreign Non Residents",
		"FR":    "Foreign Residents",
		"EAC":   "Citizens of East African Member States",
		"WR":    "Wildlife Reserve",
		"BINP":  "Bwindi Impenetrable National Park",
		"KVNP":  "Kidepo Valley National Park",
		"KNP":   "Kibale National Park",
		"LMNP":  "Lake Mburo National Park",
		"MENP":  "Mount Elgon National Park",
		"MFNP":  "Murchison Falls National Park",
		"MGNP":  "Mgahinga Gorilla National Park",
		"PAs":   "Protected Areas",
		"QENP":  "Queen Elizabeth National Park",
		"RMNP":  "Rwenzori Mountains National Park",
		"SNP":   "Semuliki National Park",
		"UWA":   "Uganda Wildlife Authority",
		"MFIFT": "Murchison Falls Invitational Fishing Tournament",
		"LTR":   "Long Term Researchers (beyond 3 years)",
		"NFSR":  "Non Foreign Student Research",
		"FSR":   "Foreign Student Research",
		"FRA/V": "Foreign Research Assistant/Volunteer",
		"UNSR":  "Ugandan Non Student Researcher",
		"UPhDS": "Ugandan PhD Student",
		"UMSCS": "Uganda MSC Students",
		"MOH":   "Ministry of Health",
		"N/A":   "Not Applicable",
	}

	// Convert data to JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal("Error converting data to JSON:", err)
	}

	// Write JSON data to a file
	jsonFile, err := os.Create("abbreviations.json")
	if err != nil {
		log.Fatal("Error creating JSON file:", err)
	}
	defer jsonFile.Close()

	_, err = jsonFile.Write(jsonData)
	if err != nil {
		log.Fatal("Error writing JSON data to file:", err)
	}

	log.Println("Conversion to JSON is complete. The output file is: abbreviations.json")

	// Convert data to CSV
	csvFile, err := os.Create("abbreviations.csv")
	if err != nil {
		log.Fatal("Error creating CSV file:", err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Write CSV header
	header := []string{"Key", "Value"}
	if err := writer.Write(header); err != nil {
		log.Fatal("Error writing CSV header:", err)
	}

	// Write CSV data
	for key, value := range data {
		record := []string{key, value}
		if err := writer.Write(record); err != nil {
			log.Fatal("Error writing CSV data:", err)
		}
	}

	log.Println("Conversion to CSV is complete. The output file is: abbreviations.csv")
}

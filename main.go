package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// https://ugandawildlife.org/wp-json/wp/v2/national-parks
// https://ugandawildlife.org/wp-json/wp/v2/national-parks/2457
// https://ugandawildlife.org/wp-json/wp/v2/media?parent=2453
// https://ugandawildlife.org/wp-json/wp/v2/animals
// https://ugandawildlife.org/book/?pid=2457

type Park struct {
	ID    int    `json:"id"`
	Link  string `json:"link"`
	Name  string `json:"name"`
	Title struct {
		Rendered string `json:"rendered"`
	} `json:"title"`
}

func main() {
	// Send an HTTP GET request to the URL
	resp, err := http.Get("https://ugandawildlife.org/wp-json/wp/v2/national-parks")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer resp.Body.Close()

	// Parse the JSON response
	var parks []Park
	err = json.NewDecoder(resp.Body).Decode(&parks)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Update park.Name to park.Title.Rendered
	for i := range parks {
		parks[i].Name = parks[i].Title.Rendered
	}

	// Save as JSON
	parksJSON, err := json.MarshalIndent(parks, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var parksWithoutTitle []map[string]interface{}
	err = json.Unmarshal(parksJSON, &parksWithoutTitle)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i := range parksWithoutTitle {
		delete(parksWithoutTitle[i], "title")
	}

	parksWithoutTitleJSON, err := json.MarshalIndent(parksWithoutTitle, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = os.WriteFile("parks.json", parksWithoutTitleJSON, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Park details saved to parks.json")

	// Save as CSV
	file, err := os.Create("parks.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header row
	header := []string{"Id", "Name", "Link"}
	err = writer.Write(header)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Write the park details
	for _, park := range parks {
		row := []string{
			fmt.Sprintf("%d", park.ID),
			park.Name,
			park.Link,
		}
		err = writer.Write(row)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	fmt.Println("Park details saved to parks.csv")
}

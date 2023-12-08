package anal

import (
	"log"
	"net/http"
)

func SendIdleRequest(hwid string) {
	url := "http://localhost:8080/analytics/new-launch"
	client := &http.Client{}
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Printf("Error creating HTTP request: %v", err)
		return
	}
	// Add headers
	request.Header.Set("HWID", hwid)
	request.Header.Set("Category", "Idle")

	// Send the request
	response, err := client.Do(request)
	if err != nil {
		log.Printf("Error sending idle request: %v", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Printf("Idle request failed with status code: %d", response.StatusCode)
		return
	}

	log.Println("Idle request sent successfully.")
}

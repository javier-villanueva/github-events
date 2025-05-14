package http

import (
	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/joho/godotenv"
	"time"
	"net/http"
	"encoding/json"
	"github/internal/models"
	"errors"
	"io"
	"os"
) 

func FetchEvents(username string) ([]models.GitHubEvent, error) {
	envErr := godotenv.Load()

	if envErr != nil {
		panic("Couldnt load env file")
	}

	timeout := 1000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	res, err := client.Get(
		"https://api.github.com/users/" + username + "/events", 
		http.Header {
			"Authorization": {"Bearer " + os.Getenv("GITHUB_APIKEY")},
		},
	)

	if (err != nil) {
		return []models.GitHubEvent{}, err
	}

	if (res.StatusCode != 200) {
		return []models.GitHubEvent{}, errors.New("No Information found") 
	}

	defer res.Body.Close()

	body, parseError := io.ReadAll(res.Body)

	if (parseError != nil) {
		return []models.GitHubEvent{}, parseError
	}

	events := make([]models.GitHubEvent, 0)
	evErr := json.Unmarshal(body, &events)

	if (evErr != nil) {
		return []models.GitHubEvent{}, evErr
	}

	return events, nil 
}


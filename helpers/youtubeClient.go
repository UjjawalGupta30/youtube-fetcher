package helpers

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var (
	apiKeys         []string
	currentKeyIndex int
	lock            sync.Mutex
)

func init() {
	apiKeys = strings.Split(os.Getenv("YOUTUBE_API_KEYS"), ",")
}

func getNextApiKey() string {
	lock.Lock()
	defer lock.Unlock()
	if currentKeyIndex >= len(apiKeys) {
		currentKeyIndex = 0
	}
	key := apiKeys[currentKeyIndex]
	currentKeyIndex++
	return key
}

func InitYouTubeClient() *youtube.Service {
	var service *youtube.Service
	var err error

	for i := 0; i < len(apiKeys); i++ {
		apiKey := getNextApiKey()
		client := &http.Client{}
		service, err = youtube.NewService(context.Background(), option.WithHTTPClient(client), option.WithAPIKey(apiKey))
		log.Printf("Initializing YouTube client with API Key %s", apiKey)
		if err == nil {
			return service
		}
		log.Printf("Failed to initialize YouTube client with API Key %s: %v", apiKey, err)
	}

	log.Fatalf("Unable to initialize YouTube client after trying all keys: %v", err)
	return nil
}

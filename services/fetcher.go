package services

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"google.golang.org/api/youtube/v3"
	"youtube-fetcher/config"
	"youtube-fetcher/helpers"
	"youtube-fetcher/models"
)

func FetchVideos() {
	service := helpers.InitYouTubeClient()
	queryInterval, err := strconv.Atoi(os.Getenv("QUERY_INTERVAL"))
	if err != nil {
		log.Printf("Error parsing queryInterval: %v", err)
		queryInterval = 10
	}
	ticker := time.NewTicker(time.Duration(queryInterval) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		videos, err := fetchYouTubeVideos(service, "official") // Assuming "official" is the search query
		if err != nil {
			log.Printf("Error fetching YouTube videos: %v", err)
			continue
		}

		if len(videos) > 0 {
			if err := storeVideos(videos); err != nil {
				log.Printf("Error storing videos in DB: %v", err)
			}
		}
	}
}

func fetchYouTubeVideos(service *youtube.Service, searchQuery string) ([]models.Video, error) {
	call := service.Search.List([]string{"id", "snippet"}).
		Q(searchQuery).
		MaxResults(50).
		Type("video").
		Order("date").
		PublishedAfter(time.Now().Add(-24 * time.Hour).Format(time.RFC3339)) // Fetch videos from the last 24 hours

	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	var videos []models.Video
	for _, item := range response.Items {
		publishTime, _ := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
		videos = append(videos, models.Video{
			ID:           item.Id.VideoId,
			Title:        item.Snippet.Title,
			Description:  item.Snippet.Description,
			ChannelTitle: item.Snippet.ChannelTitle,
			PublishTime:  publishTime,
		})
	}
	return videos, nil
}

func storeVideos(videos []models.Video) error {
	videoInterfaces := make([]interface{}, len(videos))
	for i, v := range videos {
		videoInterfaces[i] = v
	}

	_, err := config.DB.Collection("videos").InsertMany(context.Background(), videoInterfaces)
	return err
}

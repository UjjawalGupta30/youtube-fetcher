package controllers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"youtube-fetcher/config"
	"youtube-fetcher/models"
)

func GetVideos(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")

	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	var videos []models.Video
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "publish_time", Value: -1}})
	findOptions.SetSkip(int64((page - 1) * pageSize))
	findOptions.SetLimit(int64(pageSize))

	cur, err := config.DB.Collection("videos").Find(context.Background(), bson.D{}, findOptions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching videos"})
		return
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var video models.Video
		err := cur.Decode(&video)
		if err != nil {
			continue // handle error if necessary
		}
		videos = append(videos, video)
	}

	if len(videos) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No videos found"})
		return
	}

	c.JSON(http.StatusOK, videos)
}

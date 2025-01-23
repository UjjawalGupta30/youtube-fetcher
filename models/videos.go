package models

import (
	"time"
)

type Video struct {
	ID           string    `bson:"_id" json:"id"`
	Title        string    `bson:"title" json:"title"`
	Description  string    `bson:"description" json:"description"`
	ChannelTitle string    `bson:"channel_title" json:"channel_title"`
	PublishTime  time.Time `bson:"publish_time" json:"publish_time" index:"publish_time_1"`
}

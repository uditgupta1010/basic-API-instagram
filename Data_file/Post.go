package data_files

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	UserId          primitive.ObjectID `json:"userId,omitempty" bson:"userId,omitempty"`
	Id              string             `json:"id,omitempty" bson:"_id,omitempty"`
	Caption         string             `json:"caption,omitempty" bson:"caption,omitempty"`
	ImgUrl          string             `json:"imgUrl,omitempty" bson:"imgUrl,omitempty"`
	PostedTimestamp time.Time          `json:"postedTimestamp,omitempty" bson:"postedTimestamp,omitempty"`
}

type Post_j struct {
	UserId          string    `json:"userId,omitempty" bson:"userId,omitempty"`
	Id              string    `json:"id,omitempty" bson:"_id,omitempty"`
	Caption         string    `json:"caption,omitempty" bson:"caption,omitempty"`
	ImgUrl          string    `json:"imgUrl,omitempty" bson:"imgUrl,omitempty"`
	PostedTimestamp time.Time `json:"postedTimeStamp,omitempty" bson:"postedTimestamp,omitempty"`
}

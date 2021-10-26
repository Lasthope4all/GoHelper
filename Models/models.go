package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Board struct {
	Id      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title   string             `json:"title,omitempty" bson:"title,omitempty"`
	Desc    string             `json:"desc,omitempty" bson:"desc,omitempty"`
	Content string             `json:"content,omitempty" bson:"content,omitempty"`
	Tasks   []Task             `json:"tasks"`
}

type Task struct {
	Title   string `json:"title,omitempty" bson:"title,omitempty"`
	Desc    string `json:"desc,omitempty" bson:"desc,omitempty"`
	Content string `json:"content,omitempty" bson:"content,omitempty"`
}

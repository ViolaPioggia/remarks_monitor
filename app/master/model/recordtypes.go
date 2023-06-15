package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Record struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Type     string             `bson:"type,omitempty" json:"type,omitempty"`
	Content  string             `bson:"content,omitempty" json:"content,omitempty"`
	Nums     int64              `bson:"nums,omitempty" json:"nums,omitempty"`
	UpdateAt time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}

package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Remarks struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string             `bson:"username,omitempty" json:"username,omitempty"`
	Domain   string             `bson:"domain,omitempty" json:"domain,omitempty"`
	Content  string             `bson:"content,omitempty" json:"content,omitempty"`
	UpdateAt time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}

package structs

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mongoDB/internal/structs/status"
	"time"
)

type Record struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Text      string             `bson:"text" json:"text"`
	Status    status.Status      `bson:"status" json:"status"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}

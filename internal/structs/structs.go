package structs

import (
	"github.com/jafarsirojov/mongoDB/internal/structs/status"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

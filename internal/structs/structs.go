package structs

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mongoDB/internal/structs/status"
	"time"
)

type Record struct {
	ID        primitive.ObjectID `bson:"id"`
	Name      string             `bson:"name"`
	Text      string             `bson:"text"`
	Status    status.Status      `bson:"status"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

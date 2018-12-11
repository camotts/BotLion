package core

import (
	"context"
	"sync"

	"github.com/mongodb/mongo-go-driver/mongo"
)

type mongoDB struct{}

var instance *mongoDB
var instantiateOnce sync.Once

func GetMongoDB(ctx context.Context) *mongoDB {
	instantiateOnce.Do(func() {
		cl, err := mongo.Connect(ctx, "mongodb://localhost:27017")
	})
	return instance
}

func InitializeMongoDBOnce() {

}

func test() {
	cl, err := mongo.NewClient("mongodb://localhost:27017")
}

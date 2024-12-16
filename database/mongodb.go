package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var Client *mongo.Client

func InitializeMongoDB(url string) *mongo.Database {
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal("MongoDB 클라이언트 생성 실패 %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("MongoDB 연결 실패 %v", err)
	}

	Client = client

	return client.Database("cannabis_stock_db")
}

func DisconnectMongoDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if Client != nil {
		if err := Client.Disconnect(ctx); err != nil {
			log.Fatal("MongoDB 연결 종료 실패, %v", err)
		}
	}
}

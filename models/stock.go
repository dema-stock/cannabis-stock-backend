package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Stock struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Symbol   string             `bson:"symbol"`
	Name     string             `bson:"name"`
	Price    float64            `bson:"price"`
	Volume   int                `bson:"volume"`
	CreateAt time.Time          `bson:"create_at"`
	UpdateAt time.Time          `bson:"update_at"`
}

var StockCollection *mongo.Collection

// 몽고db 초기화
func InitializeStockCollection(db *mongo.Database) {
	StockCollection = db.Collection("stocks")
}

// 주식 추가
func CreateStock(ctx context.Context, stock Stock) (*mongo.InsertOneResult, error) {
	stock.CreateAt = time.Now()
	stock.UpdateAt = time.Now()
	return StockCollection.InsertOne(ctx, stock)
}

// 주식 검색
func GetStockBySymbol(ctx context.Context, symbol string) (*Stock, error) {
	var stock Stock
	err := StockCollection.FindOne(ctx, bson.M{"symbol": symbol}).Decode(&stock)
	if err != nil {
		return nil, err
	}
	return &stock, nil
}

// 모든 주식 조회
func GetAllStocks(ctx context.Context) ([]Stock, error) {
	cursor, err := StockCollection.Find(ctx, map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var stocks []Stock
	for cursor.Next(ctx) {
		var stock Stock
		if err := cursor.Decode(&stock); err != nil {
			return nil, err
		}
		stocks = append(stocks, stock)
	}

	return stocks, nil
}

// 주식 가격 업데이트
func UpdateStockPrice(ctx context.Context, symbol string, newPrice float64) error {
	update := map[string]interface{}{
		"$set": map[string]interface{}{
			"price":     newPrice,
			"update_at": time.Now(),
		},
	}
	_, err := StockCollection.UpdateOne(ctx, bson.M{"symbol": symbol}, bson.M{"$set": update})
	return err
}

// 주식 삭제
func DeleteStock(ctx context.Context, symbol string) error {
	_, err := StockCollection.DeleteOne(ctx, map[string]interface{}{"symbol": symbol})
	return err
}

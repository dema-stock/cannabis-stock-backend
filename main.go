package main

import (
	"cannabis_stock/database"
	"cannabis_stock/models"
	"cannabis_stock/services"
	"context"
	"log"
	"time"
)

func main() {
	db := database.InitializeMongoDB("mongodb://localhost:27017")
	defer database.DisconnectMongoDB()

	models.InitializeStockCollection(db)

	stockService := services.StockService{}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newStock := models.Stock{
		Symbol: "PDL",
		Name:   "실낙원",
		Price:  0,
		Volume: 0,
	}

	err := stockService.CreateStock(ctx, newStock)
	if err != nil {
		log.Fatalf("주식 추가 실패: %v", err)
	}

	log.Println("주식 추가 성공!")

	stock, err := stockService.GetStock(ctx, "PDL")
	if err != nil {
		log.Fatalf("주식 조회 실패: %v", err)
	}

	log.Printf("조회된 주식: %+v", stock)
}

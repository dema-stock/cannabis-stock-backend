package services

import (
	"cannabis_stock/models"
	"context"
)

type StockService struct{}

func (s *StockService) CreateStock(ctx context.Context, stock models.Stock) error {
	_, err := models.CreateStock(ctx, stock)
	return err
}

func (s *StockService) GetStock(ctx context.Context, symbol string) (*models.Stock, error) {
	return models.GetStockBySymbol(ctx, symbol)
}

func (s *StockService) GetAllStocks(ctx context.Context, symbol string) ([]models.Stock, error) {
	return models.GetAllStocks(ctx)
}

func (s *StockService) UpdateStock(ctx context.Context, symbol string, newPrice float64) error {
	return models.UpdateStockPrice(ctx, symbol, newPrice)
}

func (s *StockService) DeleteStock(ctx context.Context, symbol string) error {
	return models.DeleteStock(ctx, symbol)
}

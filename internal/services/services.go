package services

import (
	"context"
	"errors"
	"fmt"
	"gw-exchanger/internal/storages"
	"gw-exchanger/internal/storages/postgres"
)

var (
	ErrExchangeRatesNotAvailable = errors.New("exchange rates not available")
	ErrInvalidCurrencyPair       = errors.New("invalid currency pair")
)

// ExchangeService определяет интерфейс бизнес-логики
type ExchangeService interface {
	GetExchangeRates(ctx context.Context) (map[string]float32, error)
	GetExchangeRateForCurrency(ctx context.Context, from, to string) (float32, error)
}

// ExchangeServiceImpl реализует бизнес-логику
type ExchangeServiceImpl struct {
	storage storages.Storage
}

func NewExchangeService(storage storages.Storage) *ExchangeServiceImpl {
	return &ExchangeServiceImpl{
		storage: storage,
	}
}

func (s *ExchangeServiceImpl) GetExchangeRates(ctx context.Context) (map[string]float32, error) {
	exchange, err := s.storage.GetAllRates(ctx)
	if err != nil {
		if errors.Is(err, postgres.ErrNoRatesFound) {
			return nil, ErrExchangeRatesNotAvailable
		}
		return nil, fmt.Errorf("service: failed to get exchange rates: %w", err)
	}

	return exchange.Rates, nil
}

func (s *ExchangeServiceImpl) GetExchangeRateForCurrency(ctx context.Context, from, to string) (float32, error) {
	exchangeRate, err := s.storage.GetRateForCurrency(ctx, from, to)
	if err != nil {
		if errors.Is(err, postgres.ErrRateNotFound) {
			return 0, ErrInvalidCurrencyPair
		}
		return 0, fmt.Errorf("service: failed to get rate: %w", err)
	}

	return exchangeRate.Rate, nil
}

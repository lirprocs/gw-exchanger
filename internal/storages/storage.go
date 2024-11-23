package storages

import "context"

type Storage interface {
	GetAllRates(ctx context.Context) (*ExchangeRates, error)
	GetRateForCurrency(ctx context.Context, fromCurrency string, toCurrency string) (*ExchangeRate, error)
}

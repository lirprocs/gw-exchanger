package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gw-exchanger/internal/storages"
)

var (
	ErrNoRatesFound = errors.New("no exchange rates found")
	ErrRateNotFound = errors.New("exchange rate not found")
	ErrQueryFailed  = errors.New("query failed")
)

func (p *PostgresStorage) GetAllRates(ctx context.Context) (*storages.ExchangeRates, error) {
	const op = "postgres.GetAllRates"
	query := "SELECT from_currency, rate FROM exchange_rates WHERE to_currency = 'RUB'"

	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, ErrQueryFailed)
	}
	defer rows.Close()

	rates := make(map[string]float32)
	for rows.Next() {
		var currency string
		var rate float32
		if err := rows.Scan(&currency, &rate); err != nil {
			return nil, fmt.Errorf("%s: failed to scan row: %w", op, err)
		}
		rates[currency] = rate
	}

	if len(rates) == 0 {
		return nil, fmt.Errorf("%s: %w", op, ErrNoRatesFound)
	}

	return &storages.ExchangeRates{Rates: rates}, nil
}

func (p *PostgresStorage) GetRateForCurrency(ctx context.Context, fromCurrency, toCurrency string) (*storages.ExchangeRate, error) {
	const op = "postgres.GetRateForCurrency"
	query := "SELECT rate FROM exchange_rates WHERE from_currency = $1 AND to_currency = $2"

	var rate float32
	err := p.db.QueryRowContext(ctx, query, fromCurrency, toCurrency).Scan(&rate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, ErrRateNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, ErrQueryFailed)
	}

	return &storages.ExchangeRate{
		FromCurrency: fromCurrency,
		ToCurrency:   toCurrency,
		Rate:         rate,
	}, nil
}

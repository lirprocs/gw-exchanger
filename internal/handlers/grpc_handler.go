package handlers

import (
	"context"
	"errors"
	pb "github.com/lirprocs/proto-exchange/exchange"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gw-exchanger/internal/services"
	"log/slog"
)

type Server struct {
	pb.UnimplementedExchangeServiceServer
	log     *slog.Logger
	service services.ExchangeService
}

func Register(gRPC *grpc.Server, log *slog.Logger, service services.ExchangeService) {
	pb.RegisterExchangeServiceServer(gRPC, &Server{
		log:     log,
		service: service,
	})
}

func (s *Server) GetExchangeRates(ctx context.Context, req *pb.Empty) (*pb.ExchangeRatesResponse, error) {
	rates, err := s.service.GetExchangeRates(ctx)
	if err != nil {
		s.log.Error("failed to get exchange rates", slog.String("error", err.Error()))

		if errors.Is(err, services.ErrExchangeRatesNotAvailable) {
			return nil, status.Error(codes.NotFound, "no exchange rates available")
		}
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.ExchangeRatesResponse{Rates: rates}, nil
}

func (s *Server) GetExchangeRateForCurrency(ctx context.Context, req *pb.CurrencyRequest) (*pb.ExchangeRateResponse, error) {
	if req.FromCurrency == "" || req.ToCurrency == "" {
		return nil, status.Error(codes.InvalidArgument, "currency fields must not be empty")
	}

	rate, err := s.service.GetExchangeRateForCurrency(ctx, req.FromCurrency, req.ToCurrency)
	if err != nil {
		s.log.Error("failed to get exchange rate", slog.String("error", err.Error()))

		if errors.Is(err, services.ErrInvalidCurrencyPair) {
			return nil, status.Error(codes.NotFound, "currency pair not supported")
		}
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.ExchangeRateResponse{
		FromCurrency: req.FromCurrency,
		ToCurrency:   req.ToCurrency,
		Rate:         rate,
	}, nil
}

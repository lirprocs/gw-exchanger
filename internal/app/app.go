package app

import (
	"fmt"
	"google.golang.org/grpc"
	"gw-exchanger/internal/handlers"
	"gw-exchanger/internal/services"
	"log/slog"
	"net"
)

type App struct {
	GRPCSrv *grpc.Server
	log     *slog.Logger
	port    string
	service services.ExchangeService
}

func New(
	log *slog.Logger,
	grpcPort string,
	service services.ExchangeService,
) *App {
	gRPCServer := grpc.NewServer()
	handlers.Register(gRPCServer, log, service)

	return &App{
		GRPCSrv: gRPCServer,
		log:     log,
		port:    grpcPort,
		service: service,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "app.Run"

	log := a.log.With(slog.String("op", op))

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc server is running", slog.String("addr", lis.Addr().String()))

	if err := a.GRPCSrv.Serve(lis); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "app.Stop"

	a.log.With(slog.String("op", op)).Info("stopping gRPC server", slog.String("port", a.port))

	a.GRPCSrv.GracefulStop()
}

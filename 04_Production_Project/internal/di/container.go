package di

import (
	"database/sql"

	"github.com/Atharva0506/trading_bot/internal/config"
	delivery "github.com/Atharva0506/trading_bot/internal/delivery/http"
	"github.com/Atharva0506/trading_bot/internal/repository/postgres"
	"github.com/Atharva0506/trading_bot/internal/service"
)

type Container struct {
	UserHandler   *delivery.UserHandler
	SignalHandler *delivery.SignalHandler
	SignalService *service.SignalService
}

func NewContainer(db *sql.DB, cfg *config.Config) *Container {
	userRepo := postgres.NewUserPostgresRepo(db)
	userService := service.NewUserService(userRepo, cfg.JWT.Secret, cfg.JWT.AccessExpiry, cfg.JWT.RefreshExpiry)

	userHandler := delivery.NewUserHandler(userService)
	signalRepo := postgres.NewSignalPostgresRepo(db)
	signalService := service.NewSignalService(signalRepo)
	signalHandler := delivery.NewSignalHandler(signalService)

	return &Container{
		UserHandler:   userHandler,
		SignalHandler: signalHandler,
		SignalService: signalService,
	}
}

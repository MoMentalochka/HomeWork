package main

import (
	"context"
	"os/signal"
	"syscall"
	"time"

	"github.com/MoMentalochka/HomeWork/assembly/internal/app"
	"github.com/MoMentalochka/HomeWork/assembly/internal/config"
	"github.com/MoMentalochka/HomeWork/platform/pkg/closer"
	"github.com/MoMentalochka/HomeWork/platform/pkg/logger"
	"go.uber.org/zap"
)

const configPath = "../../deploy/compose/assembly/.env"

func main() {
	err := config.Load(configPath)
	if err != nil {
		logger.Error(context.Background(), "failed to load config: %v", zap.Error(err))
		return
	}

	appCtx, appCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer appCancel()
	defer gracefulShutdown()

	closer.Configure(syscall.SIGINT, syscall.SIGTERM)

	a, err := app.New(appCtx)
	if err != nil {
		logger.Error(appCtx, "❌ Не удалось создать приложение", zap.Error(err))
		return
	}

	err = a.Run(appCtx)
	if err != nil {
		logger.Error(appCtx, "❌ Ошибка при работе приложения", zap.Error(err))
		return
	}
}

func gracefulShutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := closer.CloseAll(ctx); err != nil {
		logger.Error(ctx, "❌ Ошибка при завершении работы", zap.Error(err))
	}
}

package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/MoMentalochka/HomeWork/order/internal/config"
	"github.com/MoMentalochka/HomeWork/platform/pkg/closer"
	"github.com/MoMentalochka/HomeWork/platform/pkg/logger"
	ordersv1 "github.com/MoMentalochka/HomeWork/shared/pkg/openapi/order/v1"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

type App struct {
	diContainer *diContainer
	orderAPI    *ordersv1.Server
	httpServer  *http.Server
	router      *chi.Mux
}

func New(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		logger.Error(ctx, "New error ", zap.Error(err))
		return nil, err
	}

	return a, nil
}

func (a *App) Run(ctx context.Context) error {
	return a.runHTTPServer(ctx)
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initDI,
		a.initLogger,
		a.initCloser,
		a.initRouter,
		a.initHTTPServer,
	}
	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initDI(_ context.Context) error {
	a.diContainer = NewDiContainer()
	return nil
}

func (a *App) initLogger(_ context.Context) error {
	return logger.Init(
		config.AppConfig().Logger.Level(),
		config.AppConfig().Logger.AsJson(),
	)
}

func (a *App) initCloser(_ context.Context) error {
	closer.SetLogger(logger.Logger())
	return nil
}

func (a *App) initRouter(_ context.Context) error {
	r := chi.NewRouter()

	// Добавляем middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))

	// Монтируем обработчики OpenAPI
	r.Mount("/", a.orderAPI)
	a.router = r

	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {

	orderAPI, err := ordersv1.NewServer(a.diContainer.OrderV1API(ctx))
	if err != nil {
		logger.Error(ctx, "ошибка создания сервера OpenAPI: %v", zap.Error(err))
	}
	a.orderAPI = orderAPI

	_ = a.initRouter(ctx)

	a.httpServer = &http.Server{
		Addr:              config.AppConfig().Http.Address(),
		Handler:           a.router,
		ReadHeaderTimeout: config.AppConfig().Http.ReadTimeOut(),
	}

	closer.AddNamed("HTTP server", func(ctx context.Context) error {
		err := a.httpServer.Shutdown(ctx)
		logger.Info(ctx, ("🛑 Завершение работы сервера..."))
		if err != nil {
			logger.Error(ctx, "❌ Ошибка при остановке сервера: %v\n", zap.Error(err))
			return err
		}
		return nil
	})

	return nil
}

func (a *App) runHTTPServer(ctx context.Context) error {
	logger.Info(ctx, fmt.Sprintf("🚀 HTTP-сервер запущен на порту %s", config.AppConfig().Http.Port()))

	err := a.httpServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error(ctx, fmt.Errorf("❌ Ошибка запуска сервера: %v", err).Error())
	}
	return nil
}

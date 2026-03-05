package di

import (
	"net/http"

	"github.com/SilverName608/go-notes/internal/api"
	"github.com/SilverName608/go-notes/internal/application"
	"github.com/SilverName608/go-notes/internal/config"
	"github.com/SilverName608/go-notes/internal/infrastructure/db"
	"github.com/SilverName608/go-notes/internal/infrastructure/repository"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Provide(config.Load),
		fx.Provide(NewPool),
		fx.Provide(repository.NewPostgresUserRepository),
		fx.Provide(repository.NewPostgresNoteRepository),
		fx.Provide(NewUserService),
		fx.Provide(application.NewNoteService),
		fx.Provide(NewMiddleware),
		fx.Provide(api.NewUserHandler),
		fx.Provide(api.NewNoteHandler),
		fx.Provide(api.NewRouter),
		fx.Invoke(RunServer),
	)
}

func NewPool(cfg *config.Config) (*pgxpool.Pool, error) {
	return db.NewPool(cfg.DBDSN)
}

func NewMiddleware(cfg *config.Config) *api.Middleware {
	return api.NewMiddleware(cfg.JWTSecret)
}

func NewUserService(repo repository.UserRepository, cfg *config.Config) *application.UserServiceImpl {
	return application.NewUserService(repo, cfg.JWTSecret)
}

func RunServer(router chi.Router, config *config.Config) {
	err := http.ListenAndServe(":"+config.HTTPPort, router)
	if err != nil {
		panic(err)
	}
}

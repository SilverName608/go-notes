package di

import (
	"fmt"
	"net/http"

	"github.com/SilverName608/go-notes/internal/api"
	"github.com/SilverName608/go-notes/internal/application"
	"github.com/SilverName608/go-notes/internal/config"
	"github.com/SilverName608/go-notes/internal/domain/service"
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
		fx.Provide(fx.Annotate(
			repository.NewPostgresUserRepository,
			fx.As(new(repository.UserRepository)),
		)),
		fx.Provide(fx.Annotate(
			repository.NewPostgresNoteRepository,
			fx.As(new(repository.NoteRepository)),
		)),
		fx.Provide(fx.Annotate(
			NewUserService,
			fx.As(new(service.UserService)),
		)),
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
	fmt.Printf("Server launch → http://localhost:%s\n", config.HTTPPort)
	err := http.ListenAndServe(":"+config.HTTPPort, router)
	if err != nil {
		panic(err)
	}
}

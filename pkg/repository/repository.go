package repository

import (
	"context"
	"live-poilot/pkg/conf"
	"live-poilot/pkg/ent"

	"entgo.io/ent/dialect/sql"

	"github.com/gofiber/fiber/v2/log"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	drv *sql.Driver
	db  *ent.Client
}

func (r *Repository) Close() error {
	return r.db.Close()
}

func NewRepository(cfg conf.AppConfig) (*Repository, func(), error) {
	var (
		drv *sql.Driver
		err error
	)
	switch cfg.Database.Driver {
	case "sqlite3", "sqlite":
		drv, err = sql.Open("sqlite3", cfg.Database.Source)
	case "pgx", "postgres":
		drv, err = sql.Open("pgx", cfg.Database.Source)
	default:
		drv, err = sql.Open(cfg.Database.Driver, cfg.Database.Source)
	}
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
		return nil, nil, err
	}
	client := ent.NewClient(
		ent.Driver(drv),
		ent.Log(func(a ...any) {
			log.Debug(a...)
		}),
	)

	if cfg.Database.Migrate {
		if err := client.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
			return nil, nil, err
		}
	}

	cleanup := func() {
		log.Info("closing the repository resources")
		if err := client.Close(); err != nil {
			log.Errorf("failed close repository resources: %v", err)
		}
	}
	return &Repository{drv: drv, db: client}, cleanup, nil
}

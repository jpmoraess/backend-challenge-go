package main

import (
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jpmoraess/backend-challenge-go/config"
	db "github.com/jpmoraess/backend-challenge-go/db/sqlc"
	"github.com/jpmoraess/backend-challenge-go/internal/application/usecase"
	"github.com/jpmoraess/backend-challenge-go/internal/infra/persistence"
	"github.com/pkg/errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var interruptsSigs = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), interruptsSigs...)
	defer stop()

	// connection pool
	pool, err := pgxpool.New(ctx, cfg.DBSource)
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}

	// db migrations
	runMigrations(cfg.MigrationURL, cfg.DBSource)

	// db store
	store := db.NewStore(pool)

	// repositories
	walletRepository := persistence.NewWalletRepositoryAdapter(store)

	// use cases
	createWallet := usecase.NewCreateWallet(walletRepository)

	err = createWallet.Execute(ctx, &usecase.CreateWalletInput{
		WalletType: "USER",
		FullName:   "John Travolta",
		Document:   "54545454",
		Email:      "john_travolta@mail.com",
		Password:   "password",
	})
	if err != nil {
		log.Fatalf("cannot create new wallet: %v", err)
	}
}

func runMigrations(sourceURL, databaseURL string) {
	m, err := migrate.New(sourceURL, databaseURL)
	if err != nil {
		log.Fatalf("cannot create new migration instance: %v", err)
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("cannot run migrations: %v", err)
	}

	fmt.Println("db migration up done")
}

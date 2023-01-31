package main

import (
	"cleanarch/api/handler"
	"cleanarch/config"
	"cleanarch/db"
	"cleanarch/repositories"
	"cleanarch/usecases"
	"database/sql"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func main() {

	initLogging()
	cfg := config.LoadConfig()

	log.Info().Msg("starting up")

	dbConn, err := sql.Open("mysql", fmt.Sprintf("%s:password@tcp(%s:3306)/rofldb", cfg.User, cfg.Host))
	if err != nil {
		log.Error().Err(err).Msg("Cannot open db connection")
	}

	err = db.RunMigrations(dbConn)
	if err != nil && err.Error() != "no change" {
		panic(err)
	}

	repo := repositories.NewUserRepository(dbConn)
	r := handler.NewGinHandler(usecases.NewUserUsecase(repo))
	r.Run(":8080")
}

func initLogging() {

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
}

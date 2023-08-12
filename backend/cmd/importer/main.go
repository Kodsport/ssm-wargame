package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sakerhetsm/ssm-wargame/internal/config"
	"github.com/sakerhetsm/ssm-wargame/internal/skolverket"
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql/driver"
	"go.uber.org/zap"
)

func main() {
	err := realMain()
	if err != nil {
		fmt.Println("error in main:", err.Error())
		os.Exit(1)
	}
}

func realMain() error {
	godotenv.Load()
	cfg, err := config.Get()
	if err != nil {
		return err
	}
	flag.Parse()

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", cfg.DB.Username, cfg.DB.Password, cfg.DB.Address, cfg.DB.Port, cfg.DB.DBName, cfg.DB.SSLMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	log, err := zap.NewDevelopment()
	if err != nil {
		return err
	}
	defer log.Sync()

	i := skolverket.New(db, log)

	if flag.Arg(0) == "skolverk" {

		return i.Import()
	} else if flag.Arg(0) == "uni" {
		return i.ImportUnis()
	} else {
		return errors.New("bad option (skolverk/uni)")
	}

}

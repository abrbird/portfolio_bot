// This is custom goose binary with sqlite3 support only.

package main

import (
	"flag"
	"github.com/abrbird/portfolio_bot/config"
	"log"
	"os"

	_ "github.com/abrbird/portfolio_bot/migrations"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir   = flags.String("dir", ".", "directory with migration files")
)

func main() {
	config_, err := config.ParseConfig("config/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	err = flags.Parse(os.Args[1:])
	if err != nil {
		log.Fatal("goose: failed on args parsing")
	}

	args := flags.Args()
	if len(args) < 1 {
		flags.Usage()
		log.Fatal("goose: not enough args")
	}

	dbstring, command := config_.Database.Uri(), args[1]

	db, err := goose.OpenDBWithDriver("postgres", dbstring)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}

	if err := goose.Run(command, db, *dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}

package main

import (
	"github.com/MMeier05/blog_aggregator/internal/config"
	"github.com/MMeier05/blog_aggregator/internal/database"
	_ "github.com/lib/pq"
	"database/sql"
	"log"
	"os"
)
type state struct {
	config *config.Config
	db *database.Queries
}

const dbURL = "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"

func main() {

	conf, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	s := &state{
		config : &conf,
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	s.db = database.New(db)

	commands := commands{
		functions : map[string]func(*state, command) error{},
	}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)

	if len(os.Args) < 2 {
		log.Fatalf("Not enough arguments!")
		os.Exit(1)
	}
	cmd := command{
		name : os.Args[1],
		args : os.Args[2:],
	}
	err = commands.run(s, cmd)
	if err != nil {
		log.Fatal(err)
	}
}

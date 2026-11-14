package main

import (
	"github.com/MMeier05/blog_aggregator/internal/config"
	"fmt"
	"os"
)
type state struct {
	config *config.Config

}

func main() {
	conf, err := config.Read()
	if err != nil {
		fmt.Print(err)
	}

	s := &state{
		config : &conf,
	}
	commands := commands{
		functions : map[string]func(*state, command) error{},
	}
	commands.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Printf("Not enough arguments\n")
		os.Exit(1)
	}
	cmd := command{
		name : os.Args[1],
		args : os.Args[2:],
	}
	err = commands.run(s, cmd)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

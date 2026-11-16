package main

import (
	"errors"
	"fmt"
	"github.com/MMeier05/blog_aggregator/internal/database"
	"github.com/google/uuid"
	"time"
	"context"
)


func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Required 'username'\n")  
	}
	user, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}
	err = s.config.SetUser(user.Name)
	if err != nil {
		return err
	}
	fmt.Printf("Username set to '%s' successfully\n", cmd.args[0])
	return nil
}


func handlerRegister(s* state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Not enough arguments. Required: 'username'\n")
	}
	if _, err := s.db.GetUser(context.Background(), cmd.args[0]); err == nil {
		return errors.New("User already exists\n")
	}
	parameters := database.CreateUserParams{
		ID : uuid.New(),
		CreatedAt : time.Now(),
		UpdatedAt : time.Now(),
		Name : cmd.args[0],
	}
	user, err := s.db.CreateUser(context.Background(), parameters)
	if err != nil {
		return err
	}
	err = s.config.SetUser(user.Name)
	if err != nil {
		return errors.New("couldn't set current user")
	}
	fmt.Printf("%s\n", user.Name)
	return nil
}

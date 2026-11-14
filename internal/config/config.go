package config

import (
	"io"
	"os"
	"encoding/json"
)
const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	userHomePath, err := os.UserHomeDir()
	if err != nil {
		return "", err 
	}
	return userHomePath + "/" + configFileName, nil
}

func write(cfg Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	byteData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, byteData, 0666) 
	if err != nil {
		return err
	}
	return nil
}

func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	file, err := os.Open(path)
	if err != nil {
		return Config{}, err 
	}
	defer file.Close()

	byteData, err := io.ReadAll(file)
	if err != nil {
		return Config{}, err
	}
	var config Config
	if err = json.Unmarshal(byteData, &config); err != nil {
		return Config{}, err
	}
	return config, nil
}

func (c Config) SetUser (userName string) error {
	c.CurrentUserName = userName
	err := write(c)
	if err != nil {
		return err
	}
	return nil
}

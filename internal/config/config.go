package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	Url      string `json:"db_url"`
	Username string `json:"current_user_name"`
}

func Read() (Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Config{}, fmt.Errorf("Error: %w", err)
	}

	file, err := os.Open(path.Join(home, configFileName))
	if err != nil {
		return Config{}, fmt.Errorf("Error: %w", err)
	}

	defer file.Close()

	config := Config{}

	data, err := io.ReadAll(file)

	if err != nil {
		return Config{}, fmt.Errorf("Error: %w", err)
	}

	if err = json.Unmarshal(data, &config); err != nil {
		return Config{}, fmt.Errorf("Error: %w", err)
	}

	return config, nil
}

func (c *Config) SetUser(user string) error {
	c.Username = user
	data, err := json.Marshal(c)

	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}

	file, err := os.Create(path.Join(home, configFileName))
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}

	defer file.Close()

	if _, err := file.Write(data); err != nil {
		return fmt.Errorf("Error: %w", err)
	}

	return nil
}

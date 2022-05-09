package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func NewConfig() *Config {
	return &Config{}
}

type Config struct {
	JSONPath  string
	ProjectID string
}

func (c *Config) Init() error {
	jsonPath, err := getJSONPath()
	if err != nil {
		return fmt.Errorf("cannot configure json path: %w", err)
	}
	projectID, err := getProjectID()
	if err != nil {
		return fmt.Errorf("cannot configure project id: %w", err)
	}

	c.JSONPath = jsonPath
	c.ProjectID = projectID
	return nil
}
func getJSONPath() (string, error) {
	err := godotenv.Load("secret.env")
	if err != nil {
		return "", err
	}
	key := "JSON_PATH"
	jsonPath := os.Getenv(key)
	if len(jsonPath) == 0 {
		return "", fmt.Errorf("getJSONPath error, invalid key")
	}
	return jsonPath, nil
}

func getProjectID() (string, error) {
	err := godotenv.Load("secret.env")
	if err != nil {
		return "", err
	}
	key := "PROJECT_ID"
	projectID := os.Getenv(key)
	if len(projectID) == 0 {
		return "", fmt.Errorf("getProjectID error, invalid key")
	}
	return projectID, nil
}

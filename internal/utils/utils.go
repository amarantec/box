package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/amarantec/box/internal"
	"github.com/joho/godotenv"
)

func findEnvFile(path string) (string, error) {
	filePath := filepath.Join(path, internal.ENVFILE)
	if _, err := os.Stat(filePath); err == nil {
		return filePath, nil
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return internal.EMPTY, err
	}

	for _, file := range files {
		if file.IsDir() {
			subDirPath := filepath.Join(path, file.Name())
			subDirFilePath, err := findEnvFile(subDirPath)
			if err == nil {
				return subDirFilePath, nil
			}
		}
	}
	return internal.EMPTY, fmt.Errorf("file .env not found in %s or anywhere", path)
}

func LoadEnv() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal("error getting actual dir")
	}

	appPath := filepath.Dir(filepath.Dir(path))
	envFile, err := findEnvFile(appPath)
	if err != nil {
		log.Fatal(".env file not found")
	}
	godotenv.Load(envFile)
}

func BuildDBConfig() (string, error) {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbPort := os.Getenv("DB_PORT")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		return internal.EMPTY, fmt.Errorf("one or more environment variables are not set")

	}

	dbConfig := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		dbHost, dbPort, dbUser, dbPassword, dbName)

	return dbConfig, nil

}

package env

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() error {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return err
	}
	return nil
}

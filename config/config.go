package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func init() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(fmt.Errorf("Error loading .env file: %w", err))
	}
}

// Config func to get env value from key
func Get(key string) string {
	return os.Getenv(key)
}

// Get config value converted to uint
func GetUint(key string) uint64 {
	val, err := strconv.ParseUint(Get(key), 10, 64)
	if err != nil {
		fmt.Println(fmt.Errorf("Error getting uit value for %s: %w", key, err))
	}
	return val
}

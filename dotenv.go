package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() {
	env := os.Getenv("AGP_ENV")
	if "" == env {
		env = "development"
	}
	godotenv.Load(".env." + env + ".local")
	if "test" != env {
		godotenv.Load(".env.local")
	}
	godotenv.Load(".env." + env)
	godotenv.Load() // The Original .env
	fmt.Println("Environment:", env)
}

// get environment variable or default value
func getenv(key string, defaultValue string) string {
	envValue := os.Getenv(key)
	if "" == envValue {
		return defaultValue
	}
	return envValue
}

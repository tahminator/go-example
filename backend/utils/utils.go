package utils

import (
	"log"
	"os"
)

// Ensures that these environment variables are defined
func ValidateEnv(envs []string) {
	for _, v := range envs {
		if os.Getenv(v) == "" {
			log.Fatalf("Failed to load %s from env file", v)
		}
	}
}

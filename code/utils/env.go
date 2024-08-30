package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv" // https://pkg.go.dev/github.com/joho/godotenv@v1.5.1
)

// load values from .env file into the system environment.
// note: .env file must be in the same dir as program entrypoint.
func Load_env_file() {
	// if there is an err loading the .env file, log it.
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

/*
Returns environment variable value that matches given key name, and whether it exists.

parameters:
  - env_key_name string: key name of the environment variable (<THIS>=<somevalue>).

returns:
  - value string: environment variable value string.
  - exists bool: whether the variable exists.
*/
func Get_env_value(env_key_name string) (string, bool) {
	value, exists := os.LookupEnv(env_key_name)
	return value, exists
}

package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv" // https://pkg.go.dev/github.com/joho/godotenv@v1.5.1
)

/*
Load values from .env file into the system environment.
note: .env file must be in the same dir as program entrypoint.

Parameters:
  - filename string: name of environment file
*/
func Load_env_file(filename string) {
	// if there is an err loading the .env file, log it.
	if err := godotenv.Load(filename); err != nil {
		log.Printf("No file %s found", filename)
	}
}

/*
Returns environment variable value that matches given key name, and whether it exists.

Parameters:
  - env_key string: key name of the environment variable (<THIS>=<somevalue>).

Returns:
  - value string: environment variable value string.
  - exists bool: whether the variable exists.
*/
func Get_env_value(env_key string) (string, bool) {
	value, exists := os.LookupEnv(env_key)
	return value, exists
}

/*
Helper function to return env value. If missing, log error and exit.

Parameters:
  - env_key string: environment variable key name.

Returns:
  - value string: value that corresponds to environment variable key name
*/
func Get_env_value_or_err(env_key string) string {
	value, exists := Get_env_value(env_key)
	if !exists {
		log.Fatalf("Environment variable key %s does not exist!", env_key)
	}
	return value
}

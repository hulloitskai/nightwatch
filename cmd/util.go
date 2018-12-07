package cmd

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: no .env file was found in the local "+
			"directory.")
	}
}

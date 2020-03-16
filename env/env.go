package env

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

const (
	// TODO: change ModeDev .env to dev.env
	DevEnvFilename  = "../.env"
	ProdEnvFilename = "../.env"
	TestEnvFilename = "../test.env"
)

func Load(filename string) {
	fmt.Printf("Loading %s env file\n", filename)
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatalf("Error loading %s file", filename)
	}
}

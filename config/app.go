package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT       string
	BucketName string
	BucketKey  string
)

func init() {
	configpath := Initenv()

	err := godotenv.Load(configpath)
	if err != nil {
		log.Fatalf("Problem loading .env file: %v", err)
		os.Exit(-1)
	}
	Port = os.Getenv("PORT")
	BucketName = os.Getenv("BUCKET_NAME")
	BucketKey = os.Getenv("BUCKET_KEY")
}

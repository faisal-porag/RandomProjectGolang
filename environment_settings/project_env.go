package environment_setting

import (
	"os"

	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string) string {
	// load .env file
	//err := godotenv.Load(".env")
	projectFilePath := os.Getenv("ENV_FILE_PATH")

	if len(projectFilePath) == 0 {
		_ = godotenv.Load(".env")
	} else {
		_ = godotenv.Load(projectFilePath + "/.env")
	}

	//if err != nil {
	//	log.Println("Error loading .env file: ", err)
	//}
	return os.Getenv(key)
}

package setting

import "github.com/joho/godotenv"

func Init() {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}
}

package pkg

import "github.com/joho/godotenv"

func LoadEnv(name string) error{
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	return nil
}
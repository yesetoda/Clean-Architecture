package main

import (
	"os"
)

func GetEnv(name string) string {
	/*/
	// info: #method 1: this is used to get the data from the .env file
		 err := godotenv.Load(".env")
		 if err != nil {
		 	log.Fatal(err)
		 }
		fmt.Println(os.Getenv("NAE"))
	*/
	// info: #method 2: directly add the enviromental var on the disk ans use it
	// info //first export it using export envVar=envValue
	// info // then just call os.Getenv("envVar")
	return os.Getenv(name)

}

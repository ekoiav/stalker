package api

import (
	"log"
	"os"

	"github.com/ekoiav/stalker/api/controllers"
	"github.com/ekoiav/stalker/api/seed"
)

var server = controllers.Server{}

func init() {
	// loads values from .env into the system
	//if err := godotenv.Load(); err != nil {
	//	log.Print("sad .env file found")
	//}
}

func Run() {

	//var err error
	//err = godotenv.Load()
	//if err != nil {
	//	log.Fatalf("Error getting env, %v", err)
	//} else {
	//	fmt.Println("We are getting the env values")
	//}

	log.Println(os.Getenv("DB_DRIVER"))
	log.Println(os.Getenv("DB_USER"))
	log.Println(os.Getenv("DB_PASSWORD"))
	log.Println(os.Getenv("DB_PORT"))
	log.Println(os.Getenv("DB_HOST"))
	log.Println(os.Getenv("DB_NAME"))

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	log.Println("oncom")

	seed.Load(server.DB)

	server.Run(":8089")

}

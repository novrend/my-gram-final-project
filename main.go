package main

import (
	"final-project/database"
	"final-project/routers"
)

func main() {
	database.StartDB()
	err := routers.StartServer().Run(":8080")
	if err != nil {
		panic(err)
	}
}

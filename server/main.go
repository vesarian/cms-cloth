package main

import (
	"github.com/vesarian/cms-cloth/database"
	"github.com/vesarian/cms-cloth/routers"
)

func main() {
	database.StartDB()
	err := routers.StartServer().Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
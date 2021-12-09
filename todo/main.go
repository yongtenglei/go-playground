package main

import (
	"log"
	"todo/dao"
	"todo/models"
	"todo/routers"
)

func main() {
	err := dao.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	dao.Db.AutoMigrate(&models.Todo{})

	r := routers.SetRouters()
	r.Run(":9030")
}

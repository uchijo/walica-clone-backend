package main

import (
	"log"

	"github.com/uchijo/walica-clone-backend/data/model"
	"github.com/uchijo/walica-clone-backend/util"
)

func init() {
	util.LoadEnv()
	util.ConnectToDB()
}

func main() {
	err := util.DB.AutoMigrate(&model.Event{}, &model.Payment{}, &model.User{})
	if err != nil {
		log.Fatal("migration failed")
	}
}

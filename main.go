package main

import (
	"ApscBlog/common/config"
	"ApscBlog/common/db"
	"ApscBlog/model"
	"ApscBlog/router"
)

func init() {
	config.ReadConfig()
	db.Connect()
	model.InitModel()
}

func main() {
	router.Run()
}

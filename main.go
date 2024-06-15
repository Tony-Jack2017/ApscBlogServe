package main

import (
	"ApscBlog/common/config"
	"ApscBlog/common/db"
	"ApscBlog/common/server"
	"ApscBlog/model"
	"ApscBlog/router"
)

func init() {
	config.ReadConfig()
	db.Connect()
	server.SetupMinio()
	model.InitModel()
}

func main() {
	router.Run()
}

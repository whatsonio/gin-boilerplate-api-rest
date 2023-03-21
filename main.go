package main

import (
	"app/config"
	"app/db"
	"app/pkg/account"
	"app/server"
)

func main() {

	config.Init()
	db.Init()

	db.GetDB().AutoMigrate(&account.Account{})
	server.Init()
}

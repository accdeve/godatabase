package main

import (
	"database/database"
	"database/service"
)

func main() {
	database.Init()
	service.ConnectMQTT()
}

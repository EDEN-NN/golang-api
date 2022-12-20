package main

import (
	"gihub.com/EDEN-NN/database"
	"gihub.com/EDEN-NN/routes"
)

func main() {
	database.Connection()
	routes.HandleRequests()
}

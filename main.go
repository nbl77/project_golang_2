package main

import (
	"Inventory_Project/config"
	"Inventory_Project/routes"
	"log"
)

func main()  {
  log.Print("Server Start at ",config.PORT)
  r := routes.Server()
  r.Start(config.PORT)
}


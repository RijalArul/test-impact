package main

import (
	"test-impact/server/databases"
	"test-impact/server/routes"
)

func main() {
	databases.StartDB()
	routes.Routes()
}

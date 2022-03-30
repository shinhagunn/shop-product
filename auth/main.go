package main

import (
	"github.com/shinhagunn/shop-auth/config"
	"github.com/shinhagunn/shop-auth/routes"
)

func main() {
	config.InitializeConfig()
	routes.InitRouter()
}

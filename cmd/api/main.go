package main

import (
	"github.com/shinhagunn/shop-product/config"
	"github.com/shinhagunn/shop-product/routes"
)

func main() {
	config.InitializeConfig()
	routes.InitRouter()
}

package main

import (
	"ecommerce/cli"
	"ecommerce/config"
)

func main() {
	config.DBConnect()
	cli.MenuLogin()

}

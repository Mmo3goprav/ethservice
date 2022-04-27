package main

import "ethservice/eth/server"

func main() {
	app := server.NewApp()
	app.Run("80")
}

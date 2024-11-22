package main

import (
	"jusbrasil-tech-challenge/cmd/api/modules"
)

func main() {
	app := modules.NewApp()
	app.Run()
}

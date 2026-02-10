package main

import (
	"learn-goravel/bootstrap"
)

func main() {
	app := bootstrap.Boot()

	app.Start()
}

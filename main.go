package main

import (
	"Unofficial_API/app"
	_ "Unofficial_API/router"
)

func main() {
	app.Instance().Run()
}

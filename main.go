package main

import (
	"Unofficial_API/app"
	_ "Unofficial_API/routers"
)

func main() {
	app.Instance().Run()
}

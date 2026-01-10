package main

import (
	"github.com/Thenecromance/Go_Blizzard_API/app"
	_ "github.com/Thenecromance/Go_Blizzard_API/routers"
)

func main() {
	app.Instance().Run()

}

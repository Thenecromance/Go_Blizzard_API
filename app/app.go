package app

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

var _Instance *App

type App struct {
	engine *gin.Engine
}

func (a *App) Run() {
	log.Println("Starting server on port 80")
	fmt.Println("Starting server on port 80")
	a.engine.Run(":80")
}

func (a *App) RegisterRoute(method, route string, handler gin.HandlerFunc) {
	a.engine.Handle(method, route, handler)
}

func Instance() *App {
	if _Instance == nil {
		_Instance = &App{
			engine: gin.Default(),
		}
	}
	return _Instance
}

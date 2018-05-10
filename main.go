package main

import (
	"os"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func main() {
	loadConfig()
	handleApp()
}

func handleApp() {
	gin.SetMode(cfg.Env)
	route := gin.New()
	route.Use(gin.RecoveryWithWriter(os.Stderr))

	svr := endless.NewServer(cfg.Address, route)
	svr.SetKeepAlivesEnabled(true)
}

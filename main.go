package main

import (
	"flag"
	"io/ioutil"
	"os"
	"strconv"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

var (
	pidFile = "log/app.pid"
)

func init() {
	os.Mkdir("log", os.ModePerm)
}

func writePid() {
	pid := syscall.Getpid()
	ioutil.WriteFile(pidFile, []byte(strconv.Itoa(pid)), 0644)
}

func main() {
	flag.Parse()

	loadConfig()
	handleApp()
}

func handleApp() {
	gin.SetMode(cfg.Env)
	route := gin.New()
	route.Use(gin.RecoveryWithWriter(os.Stderr))

	svr := endless.NewServer(cfg.Address, route)
	svr.SetKeepAlivesEnabled(true)
	svr.BeforeBegin = func(add string) {
		writePid()
	}
	// 捕获进程USR1信号，reload config
	svr.SignalHooks[endless.PRE_SIGNAL][syscall.SIGUSR1] = append(svr.SignalHooks[endless.PRE_SIGNAL][syscall.SIGUSR1], func() {
		loadConfig()
	})
	if err := svr.ListenAndServe(); err != nil {
		glog.Errorf("%s", err)
	}
}

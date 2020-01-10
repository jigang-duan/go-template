package main

import (
	"github.com/zserge/lorca"
	"time"

	"go-template/bootstrap"
	"go-template/features"
)

const (
	addr     = "127.0.0.1:9009"
	appName  = "Awesome App"
	appOwner = "djg4055108@126.com"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New(appName, appOwner)
	app.Bootstrap()
	app.Configure(features.Configure)
	return app
}

func runServer() {
	app := newApp()
	app.Listen(addr)
}

func showAndWaitWindow() {
	time.Sleep(time.Second * 3)
	webview, err := lorca.New("http://"+addr, "", 800, 600)
	if err != nil {
		panic(err)
	}
	defer webview.Close()

	// webview.SetBounds(lorca.Bounds{
	// 	WindowState: lorca.WindowStateFullscreen,
	// })

	<-webview.Done()
}

func main() {
	go runServer()
	showAndWaitWindow()
}

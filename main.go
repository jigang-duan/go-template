package main

import (
	"go-template/bootstrap"
	"go-template/features"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Awesome App", "djg4055108@126.com")
	app.Bootstrap()
	app.Configure(features.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8888")
}

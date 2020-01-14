package commands

import (
	"go-template/bootstrap"
	"go-template/features"

	"github.com/spf13/cobra"
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

func runServer() error {
	app := newApp()
	return app.Listen(addr)
}

type serverCmd struct {
	*baseCmd
}

func newServerCmd() *serverCmd {
	return &serverCmd{baseCmd: newBaseCmd(&cobra.Command{
		Use:   "server",
		Short: "启动服务",
		Long:  `启动服务`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runServer()
		},
	})}
}

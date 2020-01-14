package commands

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/zserge/lorca"
)

func showAndWaitWindow() error {
	time.Sleep(time.Second * 3)
	webview, err := lorca.New("http://"+addr, "", 800, 600)
	if err != nil {
		return err
	}
	defer webview.Close()

	// webview.SetBounds(lorca.Bounds{
	// 	WindowState: lorca.WindowStateFullscreen,
	// })

	<-webview.Done()

	return nil
}

type desktopCmd struct {
	*baseCmd
}

func newDesktopCmd() *desktopCmd {
	return &desktopCmd{baseCmd: newBaseCmd(&cobra.Command{
		Use:   "desktop",
		Short: "启动桌面应用",
		Long:  `启动桌面应用`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return showAndWaitWindow()
		},
	})}
}

package commands

import (
	"runtime"

	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

type envCmd struct {
	*baseCmd
}

func newEnvCmd() *envCmd {
	return &envCmd{baseCmd: newBaseCmd(&cobra.Command{
		Use:   "env",
		Short: "打印应用的版本和环境信息",
		Long:  `打印应用的版本和环境信息`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// printMyVersion()
			jww.FEEDBACK.Printf("GOOS=%q\n", runtime.GOOS)
			jww.FEEDBACK.Printf("GOARCH=%q\n", runtime.GOARCH)
			jww.FEEDBACK.Printf("GOVERSION=%q\n", runtime.Version())

			return nil
		},
	})}
}

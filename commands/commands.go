package commands

import (
	"fmt"
	"go-template/common/loggers"
	"go-template/config"
	"time"

	"github.com/spf13/cobra"
)

type commandsBuilder struct {
	builderCommon
	commands []cmder
}

func newCommandBuilder() *commandsBuilder {
	return &commandsBuilder{}
}

func (b *commandsBuilder) addCommands(commands ...cmder) *commandsBuilder {
	b.commands = append(b.commands, commands...)
	return b
}

func (b *commandsBuilder) addAll() *commandsBuilder {
	b.addCommands(
		newEnvCmd(),
		newServerCmd(),
		newDesktopCmd(),
	)
	return b
}

func (b *commandsBuilder) build() *rootCmd {
	h := b.newRootCmd()
	addCommands(h.getCommand(), b.commands...)
	return h
}

func addCommands(root *cobra.Command, commands ...cmder) {
	for _, command := range commands {
		cmd := command.getCommand()
		if cmd == nil {
			continue
		}
		root.AddCommand(cmd)
	}
}

type baseCmd struct {
	cmd *cobra.Command
}

func (c *baseCmd) getCommand() *cobra.Command {
	return c.cmd
}

func (c *baseCmd) flagsToConfig(cfg config.Provider) {
	initializeFlags(c.cmd, cfg)
}

func newBaseCmd(cmd *cobra.Command) *baseCmd {
	return &baseCmd{cmd: cmd}
}

type rootCmd struct {
	*baseCmd
	*commandsBuilder
}

func (r *rootCmd) getCommandsBuilder() *commandsBuilder {
	return r.commandsBuilder
}

func (b *commandsBuilder) newRootCmd() *rootCmd {
	rcmd := &rootCmd{}
	rcmd.commandsBuilder = b
	rcmd.baseCmd = &baseCmd{
		cmd: &cobra.Command{
			Use:   "app",
			Short: "根命令",
			Long:  "根命令",
			RunE: func(cmd *cobra.Command, args []string) error {
				defer b.timeTrack(time.Now(), "Total")
				go runServer()
				return showAndWaitWindow()
			},
		},
	}
	addCommands(rcmd.getCommand(), b.commands...)
	b.builderCommon.handleFlags(rcmd.cmd)
	return rcmd
}

type builderCommon struct {
	environment string

	quiet   bool
	debug   bool
	verbose bool

	logFile string
}

func (cc *builderCommon) timeTrack(start time.Time, name string) {
	if cc.quiet {
		return
	}
	elapsed := time.Since(start)
	fmt.Printf("%s in %v ms\n", name, int(1000*elapsed))
}

func (cc *builderCommon) handleCommonBuilderFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&cc.environment, "environment", "e", "", "应用环境变量")
}

func (cc *builderCommon) handleFlags(cmd *cobra.Command) {
	cc.handleCommonBuilderFlags(cmd)

}

func checkErr(logger *loggers.Logger, err error, s ...string) {
	if err == nil {
		return
	}
	if len(s) == 0 {
		logger.CRITICAL.Panicln(err)
		return
	}
	for _, message := range s {
		logger.ERROR.Panicln(message)
	}
	logger.ERROR.Panicln(err)
}

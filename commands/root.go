package commands

import (
	"fmt"
	"go-template/common/loggers"
	"go-template/config"
	"strings"

	flag "github.com/spf13/pflag"

	"github.com/spf13/cobra"
)

type Result struct {
	Code    int
	Message string
}

type Response struct {
	Result *Result

	Err error

	Cmd *cobra.Command
}

func (r Response) IsUserError() bool {
	return r.Err != nil && isUserError(r.Err)
}

func Execute(args []string) Response {
	rcmd := newCommandBuilder().addAll().build()
	cmd := rcmd.getCommand()
	cmd.SetArgs(args)

	c, err := cmd.ExecuteC()

	var resp Response

	if c == cmd && err == nil {
		resp.Result = &Result{
			Code:    0,
			Message: "OK",
		}
	}

	if err == nil {
		errCount := int(loggers.GlobalErrorCounter.Count())
		if errCount > 0 {
			err = fmt.Errorf("logged %d errors", errCount)
		}
	}

	resp.Err = err
	resp.Cmd = cmd

	return resp
}

func initializeFlags(cmd *cobra.Command, cfg config.Provider) {
	persFlagKeys := []string{
		"debug",
		"verbose",
		"logFile",
	}

	flagKeys := []string{
		"baseURL",
	}

	flagKeysForced := []string{}

	for _, key := range persFlagKeys {
		setValueFromFlag(cmd.PersistentFlags(), key, cfg, "", false)
	}

	for _, key := range flagKeys {
		setValueFromFlag(cmd.Flags(), key, cfg, "", false)
	}

	for _, key := range flagKeysForced {
		setValueFromFlag(cmd.Flags(), key, cfg, "", true)
	}
}

func setValueFromFlag(flags *flag.FlagSet, key string, cfg config.Provider, targetKey string, force bool) {
	key = strings.TrimSpace(key)
	if (force && flags.Lookup(key) != nil) || flags.Changed(key) {
		f := flags.Lookup(key)
		configKey := key
		if targetKey != "" {
			configKey = targetKey
		}

		switch f.Value.Type() {
		case "bool":
			bv, _ := flags.GetBool(key)
			cfg.Set(configKey, bv)
		case "string":
			cfg.Set(configKey, f.Value.String)
		case "stringSlice":
			bv, _ := flags.GetStringSlice(key)
			cfg.Set(configKey, bv)
		case "int":
			iv, _ := flags.GetInt(key)
			cfg.Set(configKey, iv)
		default:
			panic(fmt.Sprintf("更新判断使用 %s", f.Value.Type()))
		}
	}
}

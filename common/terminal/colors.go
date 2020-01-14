package terminal

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	isatty "github.com/mattn/go-isatty"
)

const (
	errorColor   = "\033[1;31m%s\033[0m"
	warningColor = "\033[0;33m%s\033[0m"
	noticeColor  = "\033[1;36m%s\033[0m"
)

func IsTerminal(f *os.File) bool {
	if runtime.GOOS == "windows" {
		return false
	}

	fd := f.Fd()
	return os.Getenv("TERM") != "dumb" &&
		(isatty.IsTerminal(fd) || isatty.IsCygwinTerminal(fd))
}

func Notice(s string) string {
	return colorize(s, noticeColor)
}

func Error(s string) string {
	return colorize(s, errorColor)
}

func Warning(s string) string {
	return colorize(s, warningColor)
}

func colorize(str, color string) string {
	s := fmt.Sprintf(color, doublePercent(str))
	return singlePercent(s)
}

func doublePercent(str string) string {
	return strings.Replace(str, "%", "%%", -1)
}

func singlePercent(str string) string {
	return strings.Replace(str, "%%", "%", -1)
}

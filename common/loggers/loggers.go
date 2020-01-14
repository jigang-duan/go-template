package loggers

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"time"

	jww "github.com/spf13/jwalterweatherman"

	"go-template/common/terminal"
)

// 全局错误logs计数器
var GlobalErrorCounter *jww.Counter

func init() {
	GlobalErrorCounter = &jww.Counter{}
	jww.SetLogListeners(jww.LogCounter(GlobalErrorCounter, jww.LevelError))
}

type Logger struct {
	*jww.Notepad

	Out io.Writer

	ErrorCounter *jww.Counter
	WarnCounter  *jww.Counter

	errors *bytes.Buffer
}

func (l *Logger) PrintTimerIfDelayed(start time.Time, name string) {
	elapsed := time.Since(start)
	milli := int(1000 * elapsed.Seconds())
	if milli < 500 {
		return
	}
	l.FEEDBACK.Printf("%s in %v ms", name, milli)
}

func (l *Logger) PrintTimer(start time.Time, name string) {
	elapsed := time.Since(start)
	milli := int(1000 * elapsed.Seconds())
	l.FEEDBACK.Printf("%s in %v ms", name, milli)
}

func (l *Logger) Errors() string {
	if l.errors == nil {
		return ""
	}
	return ansiColorRe.ReplaceAllLiteralString(l.errors.String(), "")
}

func (l *Logger) Reset() {
	l.ErrorCounter.Reset()
	if l.errors != nil {
		l.errors.Reset()
	}
}

func NewLogger(stdoutThreshold, logThreshold jww.Threshold, outHandle, logHandle io.Writer, saveErrors bool) *Logger {
	return newlogger(stdoutThreshold, logThreshold, outHandle, logHandle, saveErrors)
}

func NewDebugLogger() *Logger {
	return newBasicLogger(jww.LevelDebug)
}

func NewWarningLogger() *Logger {
	return newBasicLogger(jww.LevelWarn)
}

func NewErrorLogger() *Logger {
	return newBasicLogger(jww.LevelError)
}

var (
	ansiColorRe = regexp.MustCompile("(?s)\\033\\[\\d*(;\\d*)*m")
	errorRe     = regexp.MustCompile("^(ERROR|FATAL|WARN)")
)

type ansiCleaner struct {
	w io.Writer
}

func (a ansiCleaner) Write(p []byte) (n int, err error) {
	return a.w.Write(ansiColorRe.ReplaceAll(p, []byte("")))
}

type labelColorizer struct {
	w io.Writer
}

func (a labelColorizer) Write(p []byte) (n int, err error) {
	replaced := errorRe.ReplaceAllStringFunc(string(p), func(m string) string {
		switch m {
		case "ERROR", "FATAL":
			return terminal.Error(m)
		case "WARN":
			return terminal.Warning(m)
		default:
			return m
		}
	})

	_, err = a.w.Write([]byte(replaced))
	return len(p), err
}

func InitGlobalLogger(stdoutThreshold, logThreshold jww.Threshold, outHandle, logHandle io.Writer) {
	outHandle, logHandle = getLogWriters(outHandle, logHandle)

	jww.SetStdoutOutput(outHandle)
	jww.SetLogOutput(logHandle)
	jww.SetLogThreshold(logThreshold)
	jww.SetStdoutThreshold(stdoutThreshold)

}

func getLogWriters(outHandle, logHandle io.Writer) (io.Writer, io.Writer) {
	isTerm := terminal.IsTerminal(os.Stdout)
	if logHandle != ioutil.Discard && isTerm {
		logHandle = ansiCleaner{w: logHandle}
	}

	if isTerm {
		outHandle = labelColorizer{w: outHandle}
	}

	return outHandle, logHandle
}

func newlogger(stdoutThreshold, logThreshold jww.Threshold, outHandle, logHandle io.Writer, saveErrors bool) *Logger {
	errorCounter := &jww.Counter{}
	warnCounter := &jww.Counter{}
	outHandle, logHandle = getLogWriters(outHandle, logHandle)

	listeners := []jww.LogListener{
		jww.LogCounter(errorCounter, jww.LevelError),
		jww.LogCounter(warnCounter, jww.LevelWarn),
	}
	var errorBuff *bytes.Buffer
	if saveErrors {
		errorBuff = new(bytes.Buffer)
		errorCapture := func(t jww.Threshold) io.Writer {
			if t != jww.LevelError {
				return nil
			}
			return errorBuff
		}

		listeners = append(listeners, errorCapture)
	}

	return &Logger{
		Notepad:      jww.NewNotepad(stdoutThreshold, logThreshold, outHandle, logHandle, "", log.Ldate|log.Ltime, listeners...),
		Out:          outHandle,
		ErrorCounter: errorCounter,
		WarnCounter:  warnCounter,
		errors:       errorBuff,
	}
}

func newBasicLogger(t jww.Threshold) *Logger {
	return newlogger(t, jww.LevelError, os.Stdout, ioutil.Discard, false)
}

package log

import (
	"fmt"
	"log"
)

func Debug(msg string) {
	log.Println(buildMessage("debug", msg))
}

func Info(msg string) {
	log.Println(buildMessage("info", msg))
}

func Warn(msg string) {
	log.Println(buildMessage("msg", msg))
}

func Error(msg string) {
	log.Println(buildMessage("error", msg))
}

func Panic(msg string) {
	log.Panicln(buildMessage("panic", msg))
}

func Fatal(msg string) {
	log.Fatalln(buildMessage("fatal", msg))
}

func Deubgf(format string, args ...interface{}) {
	log.Println(buildMessage("debug", format, args...))
}

func Infof(format string, args ...interface{}) {
	log.Println(buildMessage("info", format, args...))
}

func Warnf(format string, args ...interface{}) {
	log.Println(buildMessage("warn", format, args...))
}

func Errorf(format string, args ...interface{}) {
	log.Println(buildMessage("error", format, args...))
}

func Panicf(format string, args ...interface{}) {
	log.Panicln(buildMessage("panic", format, args...))
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalln(buildMessage("fatal", format, args...))
}

func buildMessage(level, format string, args ...interface{}) string {
	logFmt := fmt.Sprintf("[%s] %s", level, format)
	return fmt.Sprintf(logFmt, args...)
}

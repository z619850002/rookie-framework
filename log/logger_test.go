package mylog

import (
	"testing"
	"errors"
	"fmt"
)

func TestDebug(t *testing.T) {
	info := "this is a info log"
	warn := "this is a warning"
	err := errors.New("this is a error log")

	SetLevel(InfoLevel)
	fmt.Println("level : info level")
	Info("info",info)
	Warn("warn",warn)
	Error("error", err)

	SetLevel(WarnLevel)
	fmt.Println("level : info warn")
	Info("info",info)
	Warn("warn",warn)
	Error("error", err)

	SetLevel(ErrorLevel)
	fmt.Println("level : error level")
	Info("info",info)
	Warn("warn",warn)
	Error("error", err)

	SetLevel(FatalError)
	fmt.Println("level : fatal level")
	Info("info",info)
	Warn("warn",warn)
	Error("error", err)
	//Fatal("%v", err)

	SetLevel(10)
	fmt.Println("level : default level")
	Info("info",info)
	Warn("warn",warn)
	Error("error", err)
}

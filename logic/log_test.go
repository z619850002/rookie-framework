package logic

import (
	"testing"
	"fmt"
)


func TestNewLogicLogDecorator(t *testing.T) {
	fmt.Println("Process with log.")
	handler := NewHandler()
	handler.WrapLog()
	//Run the process.
	handler.Operate1()
}
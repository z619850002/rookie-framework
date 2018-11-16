package logic

import (
	"testing"
	"fmt"
)

func TestNewLogicMemoryDecorator(t *testing.T) {
	fmt.Println("Process with memory.")
	handler := NewHandler()
	handler.WrapMemory()
	//Run the process.
	handler.Operate1()
}
package logic

import (
	"fmt"
	"testing"
)

//The handler just like the handler in the server.
type Handler struct {
	logicModule LogicBlock
}

func NewHandler() *Handler{
	return &Handler{logicModule:NewBaseLogicBlock()}
}

//Add log system to this handler.
func (h *Handler)WrapLog(){
	h.logicModule = NewLogWrapper(h.logicModule)
}

//Add memory calculation system to this handler.
func (h *Handler)WrapMemory(){
	h.logicModule = NewMemoryWrapper(h.logicModule)
}

func (h *Handler) Operate1(){
	res , err := h.logicModule.Call1(Add, 1.0 , 2.3)
	if (err != nil){
		fmt.Println(err)
	}
	fmt.Println(res)
}

//The function can be written outside.
func Add(args... interface{})(interface{} , error){
	if (len(args) !=2){
		fmt.Println(len(args))
		return -1 , nil
	}
	i1 := args[0].(float64)
	i2 := args[1].(float64)
	return i1+i2 , nil
}

func TestNewBaseLogicBlock(t *testing.T) {
	fmt.Println("Basic process")
	handler := NewHandler()
	//Run the process.
	handler.Operate1()
}
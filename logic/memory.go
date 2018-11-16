package logic

import (
		"unsafe"
	"hub000.xindong.com/rookie/rookie-framework/log"
)

//TODO:This is just a demo, need to be modified in the future.
//LogicMemoryDecorator is one decorator that can log the size of the input.
type MemoryWrapper struct {
	block LogicBlock
}

//NewLogicMemoryDecorator will create a new logic memory decorator.
func NewMemoryWrapper(block LogicBlock)MemoryWrapper{
	decorator := MemoryWrapper{block:block}
	return decorator
}


func (h MemoryWrapper) Call0(f func(args... interface{}) error ,args... interface{} ) (err error){
	mylog.Info("Input size is " , unsafe.Sizeof(args))
	err = h.block.Call0(f , args...)
	return
}




func (h MemoryWrapper) Call1(f func(args... interface{})(interface{} , error) ,args... interface{} ) (ret interface{} , err error){
	mylog.Info("Input size is " , unsafe.Sizeof(args))
	ret , err = h.block.Call1(f , args...)
	return
}


func (h MemoryWrapper) CallN(f func(args... interface{}) ([]interface{} , error) ,args... interface{} ) (ret []interface{} , err error){
	mylog.Info("Input size is " , unsafe.Sizeof(args))
	ret , err = h.block.CallN(f , args...)
	return
}

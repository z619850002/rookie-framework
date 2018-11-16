package logic

import (
	"hub000.xindong.com/rookie/rookie-framework/log"
)

//TODO:This is just a demo, need to be modified in the future.
//LogicLogDecorator is one decorator that can log the input, output
// and something else of a logic module.
type LogWrapper struct {
	block 	LogicBlock
}

//NewLogicLogWrapper takes in a logic module and wraps it by this decorator so that
// that module will log something when it`s running.
func NewLogWrapper(module LogicBlock)LogicBlock{
	decorator := &LogWrapper{block:module}
	return decorator
}


func (h LogWrapper) Call0(f func(args... interface{}) error ,args... interface{} ) (err error){
	mylog.Info("Input is " , args)
	err = h.block.Call0(f , args...)
	mylog.Info("Error is" , err)
	return
}




func (h LogWrapper) Call1(f func(args... interface{})(interface{} , error) ,args... interface{} ) (ret interface{} , err error){
	mylog.Info("Input is " , args)
	ret , err = h.block.Call1(f , args...)
	mylog.Info("Result is " , ret)
	mylog.Info("Error is" , err)
	return
}




func (h LogWrapper) CallN(f func(args... interface{}) ([]interface{} , error) ,args... interface{} ) (ret []interface{} , err error){
	mylog.Info("Input is " , args)
	ret , err = h.block.CallN(f , args...)
	mylog.Info("Result is " , ret)
	mylog.Info("Error is" , err)
	return
}
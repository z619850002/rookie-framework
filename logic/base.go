package logic

//BaseLogicBlock is one general implementation of the logic block.
type BaseLogicBlock struct {

}

//NewBaseLogicBlock will create a new base logic module and return it.
func NewBaseLogicBlock()LogicBlock{
	return BaseLogicBlock{}
}

//Call0 takes in a function returns just error and parameters.
func (h BaseLogicBlock) Call0(f func(args... interface{}) error ,args... interface{} ) (err error){
	err = f(args...)
	return
}

//Call1 takes in a function that returns one return value and the error and parameters.
func (h BaseLogicBlock) Call1(f func(args... interface{})(interface{} , error) ,args... interface{} ) (ret interface{} , err error){
	ret , err = f(args...)
	return
}

//CallN takes in a function that returns a list of return value and the error.
func (h BaseLogicBlock) CallN(f func(args... interface{}) ([]interface{} , error) ,args... interface{} ) (ret []interface{} , err error){
	ret , err = f(args...)
	return
}

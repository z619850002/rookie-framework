package logic


//LogicBlock is the base interface of all logic process.
type LogicBlock interface {
	//Call0 has no return value except for error. It takes in a function and the parameters, then
	//operate it and return the error.
	Call0(f func(args... interface{}) error ,args... interface{} ) (err error)
	//Call1 has one return value.
	Call1(f func(args... interface{})(interface{} , error) ,args... interface{} ) (ret interface{} , err error)
	//CallN has a list of return value.
	CallN(f func(args... interface{}) ([]interface{} , error) ,args... interface{} ) (ret []interface{} , err error)
}


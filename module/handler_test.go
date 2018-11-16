package module

import (
	"fmt"
	"testing"
	"hub000.xindong.com/rookie/rookie-framework/communication"
	"strconv"
	"sync"
	"time"
)

type TestBuf struct {
	message 	string
}

func (h TestBuf) GetProtocolNum() int{
	return -500
}


//BaseHandler is a basic implementation of handler.
type BaseHandler struct {
	module *Module
}


func (h *BaseHandler) Handle(message *communication.MessageWrapper){
	time.Sleep(time.Second)
	fmt.Println(message.GetMessage())
	message.AddResponse(TestBuf{message:"response!"})
}



func TestHandler(t *testing.T) {
	//Module1
	module := NewModule()
	communicator := communication.NewRPCCommunicator()
	handler := &BaseHandler{}

	module.SetCommunicator(communicator)
	module.SetHandler(handler)


	//Module2
	module2 := NewModule()
	communicator2 := communication.NewRPCCommunicator()
	handler2 := &BaseHandler{}

	module2.SetCommunicator(communicator2)
	module2.SetHandler(handler2)

	//Run 2 modules
	module.Run()
	module2.Run()


	var wg sync.WaitGroup
	for i := 0;i<5;i++{
		wg.Add(1)
		ii := i
		go func() {
			//Send messages.
			_ , err := module.Call(module2.Communicator , &TestBuf{message:"hello!" + strconv.Itoa(ii)})
			if (err == nil){
				t.Error(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
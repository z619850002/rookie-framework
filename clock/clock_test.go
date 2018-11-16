package clock

import (
	"testing"
	"time"
	"fmt"
	"hub000.xindong.com/rookie/rookie-framework/module"
	"hub000.xindong.com/rookie/rookie-framework/communication"
	"hub000.xindong.com/rookie/rookie-framework/protobuf"
	"sync"
)

type TestBuf struct {
	message 	string
}

func (h TestBuf) GetProtocolNum() int{
	return -500
}


//BaseHandler is a basic implementation of handler.
type BaseHandler struct {
	module *module.Module
}


func (h *BaseHandler) Handle(message *communication.MessageWrapper){
	time.Sleep(time.Second)
	fmt.Println(message.GetMessage())
	message.AddResponse(TestBuf{message:"response!"})
}




func TestClock_SendDelayed(t *testing.T) {
	module := module.NewModule()
	communicator := communication.NewRPCCommunicator()
	module.SetCommunicator(communicator)
	module.SetHandler(&BaseHandler{module:module})
	module.Run()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	clock := NewClock()
	clock.SendDelayed(communicator , TestBuf{} , func(retChan chan protobuf.Buf) {
		res,ok := <- retChan
		if ok{
			fmt.Println(res)
		}
		wg.Done()
	} ,time.Second * 3)
	wg.Wait()
}

func TestClock_SendContinuously(t *testing.T) {
	module := module.NewModule()
	communicator := communication.NewRPCCommunicator()
	module.SetCommunicator(communicator)
	module.SetHandler(&BaseHandler{module:module})
	module.Run()

	wg := &sync.WaitGroup{}
	wg.Add(3)
	clock := NewClock()
	clock.SendContinuously(communicator , TestBuf{} , func(retChan chan protobuf.Buf) {
		res,ok := <- retChan
		if ok{
			fmt.Println(res)
		}
		wg.Done()
	} ,time.Second * 1)
	wg.Wait()
}

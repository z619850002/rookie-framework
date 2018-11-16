package communication

import (
	"testing"
	"fmt"
	"time"
	"hub000.xindong.com/rookie/rookie-framework/protobuf"
)

type TestBuf2 struct {
	message 	string
}

func (h *TestBuf2) GetProtocolNum() int{
	return -500
}


func TestMultiRPCCommunicator_Call(t *testing.T) {
	communicator1 := NewMultiRPCCommunicator(10).Run()
	communicator2 := NewMultiRPCCommunicator(10).Run()
	for i := 0;i<1000;i++{
		response , err := communicator1.Call(communicator2 , nil)
		if (err != nil){
			fmt.Println(err)
		}else {
			fmt.Println(i,":",response)
		}
	}
}


func TestMultiRPCCommunicator_GetInputChan(t *testing.T) {
	communicator1 := NewMultiRPCCommunicator(10)
	retChan := make(chan protobuf.Buf)
	m := NewMessageWrapper(&TestBuf2{""} , retChan)
	for i:=0;i<3;i++{
		go func() {
			time.Sleep(time.Second*3)
			communicator1.getFreeChannel() <- &m
		}()
	}
	go func() {
		time.Sleep(time.Second*6)
		communicator1.getFreeChannel() <- &m
	}()
	for i:=0;i<4;i++{
		fmt.Println("Blocked!")
		<-communicator1.GetInputChan()
		fmt.Println("Get response!")
	}
}
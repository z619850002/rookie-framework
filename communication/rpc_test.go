package communication

import (
	"testing"
		"time"
	"fmt"
	"hub000.xindong.com/rookie/rookie-framework/protobuf"
	"sync"
)

type TestBuf struct {
	message 	string
}

func (h *TestBuf) GetProtocolNum() int{
	return -500
}


func TestRPCCommunicator_Call(t *testing.T) {
	//communicator1 := NewRPCCommunicator().Run()
	//communicator2 := NewRPCCommunicator().Run()
	//for i := 0;i<1000;i++{
	//	response , err := communicator1.Call(communicator2 , nil)
	//	if (err != nil){
	//		fmt.Println(err)
	//	}else {
	//		fmt.Println(i,":",response)
	//	}
	//}
	communicator3 := NewRPCCommunicator().Run()
	communicator3.lock = false
	communicator3.MaxWaittingTime = 3*time.Second
	resp,err := communicator3.Call(communicator3 , nil)
	if (err != nil){
		t.Error(err)
	}
	fmt.Println(resp)
}


func TestRPCCommunicator_GetInput(t *testing.T) {
	communicator1 := NewRPCCommunicator()
	retChan := make(chan protobuf.Buf)
	m := NewMessageWrapper(&TestBuf{""} , retChan)
	for i:=0;i<3;i++{
		go func() {
			time.Sleep(time.Second*3)
			communicator1.inputChan <- &m
		}()
	}
	go func() {
		time.Sleep(time.Second*6)
		communicator1.inputChan <- &m
	}()
	for i:=0;i<4;i++{
		fmt.Println("Blocked!")
		<-communicator1.GetInputChan()
		fmt.Println("Get response!")
	}
}

func BenchmarkRPCCommunicator_CallAsync(b *testing.B) {
	communicator1 := NewRPCCommunicator().Run()
	communicator2 := NewRPCCommunicator().Run()
	var wg 	sync.WaitGroup
	wg.Add(1000)
	for i := 0;i<1000;i++{
		localWG := &wg
		message := TestBuf{"message"}
		err := communicator1.CallAsync(communicator2 , &message , func(retChan chan protobuf.Buf) {
			localWG.Done()
			response := <- retChan
			fmt.Println(response)
		})
		if (err != nil){
			fmt.Println(err)
		}
	}
	wg.Wait()
}
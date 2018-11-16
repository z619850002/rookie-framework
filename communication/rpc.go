package communication

import (
	"fmt"
	"github.com/pkg/errors"
	"sync"
	"hub000.xindong.com/rookie/rookie-framework/protobuf"
	"hub000.xindong.com/rookie/rookie-framework/syserr"
	"hub000.xindong.com/rookie/rookie-framework/log"
	"time"
)

//RPCComminicator is one kind of communicator that imitates RPC protocol.
type RPCCommunicator struct {
	MaxWaittingTime time.Duration
	inputChan	chan *MessageWrapper
	lock 		bool
	mutex 		*sync.Mutex
}

//NewRPCCommunicator will create a new RPCCommunicator and return it back.
func NewRPCCommunicator() *RPCCommunicator{
	return &RPCCommunicator{
		inputChan:make(chan *MessageWrapper),
		mutex:&sync.Mutex{},
		MaxWaittingTime:time.Second * 10,
		lock:true,
	}
}

//Call is sync
func (h * RPCCommunicator) Call(communicator Communicator , message protobuf.Buf) (protobuf.Buf , error){
	//The channel to get response
	retChan := make(chan protobuf.Buf)

	//Create a wrapper.
	messageWrapper := NewMessageWrapper(message , retChan)



	//Send message to the target.
	communicator.Send(&messageWrapper)


	//Get the time. The method can`t wait for more than the max limitation.
	loop:
	for {
		//Get the response from the channel.
		select {
			case response , ok := <- retChan:{
				
				if ok {
					var err error

					if (response == nil) {
						return nil, errors.New("It is empty")
						//Check the system error protocol in this level.
					}else if (response.GetProtocolNum() == protobuf.ErrorNum){
						err = syserr.RequestError{}
					}else if (response.GetProtocolNum() == protobuf.EmptyRespNum){
						err = syserr.EmptyMessageError{}
					}
					return response , err
				}
			}
			case <-time.After(h.MaxWaittingTime):{
				if h.MaxWaittingTime>0{
					break loop
				}
			}
		}
	}
	err := syserr.NoResponseError{}
	mylog.Error("Bad response for the request: ",err)
	return nil , err
}

//CallAsync is an async method.
func (h *RPCCommunicator) CallAsync(communicator Communicator , message protobuf.Buf , callback func(retChan chan protobuf.Buf)) (err error){
	//The channel to get response
	retChan 	:= make(chan protobuf.Buf)
	//Create a wrapper.
	messageWrapper := NewMessageWrapper(message , retChan)
	//Send message to the target.
	communicator.Send(&messageWrapper)
	//This is async so the method below need to run in another thread.
	go func() {
		callback(retChan)
	}()
	//There is no error now.
	return
}




func (h * RPCCommunicator) Send(wrapper *MessageWrapper){
	defer func() {
		if err:=recover();err!=nil{
			mylog.Error("The response channel has closed.",err) // log the error
		}
	}()
	//If this block is locked, other modules can`t send message to this module anymore.
	if (!h.lock) {
		go func() {
			//TODO:This is just a demo.
			h.inputChan <- wrapper

		}()
	}else{
		//This must run in another thread!!!!!!!!!!!!!!!
		go func() {

			//TODO: add module name here.
			mylog.Error("The module is locked!")
			//TODO:This channel may close firstly.
			wrapper.AddResponse(protobuf.ErrorResp{})
		}()
	}
}

//GetInput will block until some input was send to the input channel.
func (h * RPCCommunicator) GetInputChan() (chan *MessageWrapper){
	return h.inputChan
}

//Run just use for test.
func (h * RPCCommunicator) Run() *RPCCommunicator{
	h.lock = false
	go func() {
		for ;;{
			select {
			case m := <-h.inputChan :{
				fmt.Println(m.GetMessage())
				m.AddResponse(nil)

			}
			}
		}
	}()
	return h
}



func (h * RPCCommunicator) Lock(){
	h.mutex.Lock()
	h.lock = true
	//TODO:Clear the message in the channel. Maybe something more is better.
	h.inputChan = make(chan *MessageWrapper)
	h.mutex.Unlock()
}

func (h * RPCCommunicator) UnLock(){
	h.mutex.Lock()
	h.lock = false
	h.mutex.Unlock()
}


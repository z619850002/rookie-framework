package communication


import (
"sync"
"hub000.xindong.com/rookie/rookie-framework/protobuf"
"hub000.xindong.com/rookie/rookie-framework/syserr"
"hub000.xindong.com/rookie/rookie-framework/log"
	"fmt"
)

//MultiRPCComminicator is one kind of communicator that imitates RPC protocol. It has multiple channels
// so that its speed is higher than the communicator.
type MultiRPCCommunicator struct {
	inputChans	[]chan *MessageWrapper
	chanCondition	[]bool
	lock 		bool
	mutex 		*sync.Mutex
}

//NewRPCCommunicator will create a new RPCCommunicator and return it back.
func NewMultiRPCCommunicator(size int) *MultiRPCCommunicator{
	comm := &MultiRPCCommunicator{
		mutex:&sync.Mutex{},
	}
	//Add channels to the buffer.
	for i:=0;i<size;i++{
		comm.inputChans = append(comm.inputChans, make(chan *MessageWrapper))
		comm.chanCondition = append(comm.chanCondition, false)
	}
	return comm
}

//Call is sync
func (h * MultiRPCCommunicator) Call(communicator Communicator , message protobuf.Buf) (protobuf.Buf , error){
	//The channel to get response
	retChan := make(chan protobuf.Buf)

	//Create a wrapper.
	messageWrapper := NewMessageWrapper(message , retChan)
	//Send message to the target.
	communicator.Send(&messageWrapper)
	for {
		select {
		case response , ok := <- retChan:{
			if ok {
				return response , nil
			}
		}
		}
	}
	err := syserr.NoResponseError{}
	mylog.Error("Bad response for the request: ",err)
	return nil , err
}

//CallAsync is an async method.
func (h *MultiRPCCommunicator) CallAsync(communicator Communicator , message protobuf.Buf , callback func(retChan chan protobuf.Buf)) (err error){
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




func (h * MultiRPCCommunicator) Send(wrapper *MessageWrapper){
	defer func() {
		if err:=recover();err!=nil{
			mylog.Error("The response channel has closed.",err) // log the error
		}
	}()
	//If this block is locked, other modules can`t send message to this module anymore.
	if (!h.lock) {
		//TODO:This is just a demo.
		inputChan := h.getFreeChannel()
		inputChan <- wrapper
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
func (h * MultiRPCCommunicator) GetInputChan() (chan *MessageWrapper){
	channel := h.GetUsedChannel()
	return channel
}

//Run just use for test.
func (h * MultiRPCCommunicator) Run() *MultiRPCCommunicator{
	go func() {
		for ;;{
			select {
			case m := <-h.GetInputChan() :{
				fmt.Println(m.GetMessage())
				m.AddResponse(nil)
			}
			}
		}
	}()
	return h
}

func (h * MultiRPCCommunicator) Lock(){
	h.mutex.Lock()
	h.lock = true
	//TODO:Clear the message in the channel. Maybe something more is better.
	for i:=0;i<len(h.inputChans);i++{
		h.chanCondition[i] = false
		h.inputChans[i] = make(chan *MessageWrapper)
	}
	h.mutex.Unlock()
}

func (h * MultiRPCCommunicator) UnLock(){
	h.mutex.Lock()
	h.lock = false
	h.mutex.Unlock()
}

//Get the free channel.
func (h *MultiRPCCommunicator)getFreeChannel() (chan *MessageWrapper){
	for ; ;  {
		//If some channel is free.
		for i:=0;i<len(h.inputChans);i++{
			if !h.chanCondition[i]{
				h.chanCondition[i] = true
				return h.inputChans[i]
			}
		}
	}
	////create a new channel.
	//h.mutex.Lock()
	//defer func() {
	//	h.mutex.Unlock()
	//}()
	//newChan := make(chan *MessageWrapper)
	//h.inputChans = append(h.inputChans, newChan)
	//h.chanCondition = append(h.chanCondition , true)
	//return newChan
}

func (h *MultiRPCCommunicator) GetUsedChannel()(chan *MessageWrapper){
	for ;;{
		//If some channel is free.
		for i:=0;i<len(h.inputChans);i++{
			if h.chanCondition[i]{
				h.chanCondition[i] = false
				return h.inputChans[i]
			}
		}
	}
}
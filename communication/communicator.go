package communication

import "hub000.xindong.com/rookie/rookie-framework/protobuf"

type Communicator interface {
	//This is sync.
	Call(communicator Communicator , message protobuf.Buf) (protobuf.Buf , error)
	//This is async.
	CallAsync(communicator Communicator , message protobuf.Buf , callback func(retChan chan protobuf.Buf)) (err error)
	//Run will run the main thread of this communicator.
	Send(message *MessageWrapper)
	GetInputChan() chan *MessageWrapper
	//Lock the communicator. It won`t accept any message.
	Lock()
	//UnLock the communicator. It will begin to accept message.
	UnLock()
}

func NewTestCommunication() *Communicator{
	return nil
}


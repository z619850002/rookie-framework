package clock

import (
	"hub000.xindong.com/rookie/rookie-framework/communication"
	"hub000.xindong.com/rookie/rookie-framework/protobuf"
	"time"
	"hub000.xindong.com/rookie/rookie-framework/log"
)

//The Clock is an object that can send message to some module under some limitations.
type Clock struct {
	Communicator communication.Communicator
}

func NewClock() *Clock{
	return &Clock{
		Communicator:communication.NewRPCCommunicator(),
	}
}


func (h *Clock) SendDelayed(communicator communication.Communicator , message protobuf.Buf, callback func(retChan chan protobuf.Buf),delay time.Duration){
	time.AfterFunc(delay , func() {
		err := h.Communicator.CallAsync(communicator , message , callback)
		//FIXME: There is no error now.
		if (err != nil){
			mylog.Error("Error in the Clock part, " ,err)
		}
	})
}

//SendContinuously takes in the communicator, message and the interval. It will send the message
// to the communicator once in a while.
func (h *Clock) SendContinuously(communicator communication.Communicator , message protobuf.Buf, callback func(retChan chan protobuf.Buf), interval time.Duration){
	ticker := time.NewTicker(interval)
	go func() {
		for _ = range ticker.C{
			err := h.Communicator.CallAsync(communicator , message , callback)
			//FIXME: There is no error now.
			if (err != nil){
				mylog.Error("Error in the Clock part, " ,err)
			}
		}
	}()
}
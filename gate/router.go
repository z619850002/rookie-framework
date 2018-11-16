package gate

import (
	"hub000.xindong.com/rookie/rookie-framework/communication"
	"hub000.xindong.com/rookie/rookie-framework/protobuf"
	"hub000.xindong.com/rookie/rookie-framework/log"
)

type Router struct {
	p       *Processor
	route 	map[int] communication.Communicator
	Communicator	communication.Communicator
}

func NewRouter() *Router{
	router := Router{}
	router.route = make(map[int] communication.Communicator)
	router.Communicator = communication.NewRPCCommunicator()
	return  &router
}

func (h *Router) Operate(requset protobuf.CSRequest){
	mylog.Info("Receive a request. " , requset.Message)
	go func() {
		targetCommunicator,ok := h.route[requset.MessageProtocolNum]
		if ok{
			resp , err := h.Communicator.Call(targetCommunicator , requset)
			if (err != nil){
				mylog.Error("Error when doing some operations: %v" , err)
			}
			if resp.GetProtocolNum() == protobuf.ResponseNum {
				h.p.Send(resp.(protobuf.CSResponse))
			}else{
				mylog.Error("incorrect protocol number. ", resp.GetProtocolNum())
			}
		}else {

		}
	}()
}


func (h *Router) Regist(id int , comm 	communication.Communicator){
	if _,ok := h.route[id];ok{
		mylog.Warn("Repeatable route")
	}
	h.route[id] = comm
}

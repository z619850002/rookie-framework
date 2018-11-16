package communication

import (
	"hub000.xindong.com/rookie/rookie-framework/protobuf"
	"hub000.xindong.com/rookie/rookie-framework/log"
)

//MessageWrapper is a wrapper that provides message more attributes.
type MessageWrapper struct {
	message 	protobuf.Buf
	retChan		chan protobuf.Buf
}

func NewMessageWrapper(message protobuf.Buf , retChan chan protobuf.Buf) MessageWrapper{
	return MessageWrapper{
		message:message,
		retChan:retChan,
	}
}

//GetMessage returns the message content in this wrapper.
func (h *MessageWrapper) GetMessage() protobuf.Buf{
	return h.message
}

//AddResponse takes in the response and send it to the retChan in this wrapper.
func (h *MessageWrapper) AddResponse(reponse protobuf.Buf){
	defer func() {
		if err:=recover();err!=nil{
			mylog.Error("The response channel has closed.",err) // log the error
		}
	}()
	h.retChan <- reponse
}



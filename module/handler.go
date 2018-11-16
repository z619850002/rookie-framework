package module

import (
	"hub000.xindong.com/rookie/rookie-framework/communication"
)

//Handler is the message handler that handler message.
type Handler interface {
	//TODO: This need to be modified in the future.
	Handle(message *communication.MessageWrapper)
}


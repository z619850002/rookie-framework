package gate

import (
	"hub000.xindong.com/rookie/rookie-framework/protobuf"
	"hub000.xindong.com/rookie/rookie-framework/log"
)

type MsgHandler interface {
	Unmarshal(int, []byte, string) (protobuf.CSRequest, error)
	Marshal(resp protobuf.CSResponse) (*Message, error)
}

type Processor struct {
	g  			*WSGate
	Router		*Router
	handler 	MsgHandler
}

func NewProcessor(h MsgHandler) *Processor{
	router := NewRouter()
	processor := &Processor{}
	router.p = processor
	processor.Router = router
	processor.handler = h
	return processor
}

func (p *Processor)MessageHandler(conn *WsConn, t int ,msg []byte){
	req, err := p.handler.Unmarshal(t, msg ,conn.player)
	if err != nil {
		mylog.Error("MessageHandler fail.",err)
	}else{
		p.Router.Operate(req)
	}
}

func (p *Processor)Send(resp protobuf.CSResponse){
	message, err := p.handler.Marshal(resp)
	if err != nil {
		mylog.Error("Send fail.",err)
	}else{
		p.g.SendToPlayer(resp.ClientID, message)
	}
}







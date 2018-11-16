package db

import (
	"hub000.xindong.com/rookie/rookie-framework/communication"
	"hub000.xindong.com/rookie/rookie-framework/protobuf"
	"hub000.xindong.com/rookie/rookie-framework/syserr"
)

type Handler struct {
	module 	*DBModule
}

func (h *Handler) Handle(message *communication.MessageWrapper){
	buf := message.GetMessage()
	switch buf.GetProtocolNum() {
	case protobuf.DBExecReqNum:
		//Excel request
		req:=buf.(protobuf.DBExecReq)
		err:=h.module.Exec(req.DsName,req.StrSQL,req.Args...)
		resp:=protobuf.DBExecResp{Status:protobuf.Status{Error:err}}
		message.AddResponse(resp)
	case protobuf.DBQueryReqNum:
		//Query request
		req:=buf.(protobuf.DBQueryReq)
		rows,err:=h.module.Query(req.DsName,req.StrSQL,req.Args...)
		resp:=protobuf.DBQueryResp{Rows:rows,Status:protobuf.Status{Error:err}}
		message.AddResponse(resp)
	case protobuf.DBQueryRowReqNum:
		//Query request
		req:=buf.(protobuf.DBQueryRowReq)
		row,err:=h.module.QueryRow(req.DsName,req.StrSQL,req.Args...)
		resp:=protobuf.DBQueryRowResp{Row:row,Status:protobuf.Status{Error:err}}
		message.AddResponse(resp)
	default:
		message.AddResponse(protobuf.ErrorResp{Status:protobuf.Status{Error:syserr.RequestError{}}})
	}
}

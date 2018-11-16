package cache

import (
	"hub000.xindong.com/rookie/rookie-framework/communication"
	"hub000.xindong.com/rookie/rookie-framework/protobuf"
	"hub000.xindong.com/rookie/rookie-framework/syserr"
	"hub000.xindong.com/rookie/rookie-framework/log"
)


type Handler struct {
	module 	*CacheModule
}


func (h *Handler) Handle(message *communication.MessageWrapper){
	buf := message.GetMessage()
	switch buf.GetProtocolNum() {
	case protobuf.CacheGetReqNum:{
		//GetReq
		req := buf.(protobuf.CacheGetReq)
		value , err := h.module.Get(req.Name , req.Key)
		resp := protobuf.CacheGetResp{Value:value , Status:protobuf.Status{Error:err}}
		message.AddResponse(resp)
	}

	case protobuf.CacheSetReqNum:{
		//SetReq
		req := buf.(protobuf.CacheSetReq)
		err := h.module.Set(req.Name , req.Key , req.Value , req.ExpireTime)
		resp := protobuf.CacheSetResp{Status:protobuf.Status{Error:err}}
		message.AddResponse(resp)
	}

	case protobuf.CacheDeleteReqNum:{
		//DeleteReq
		req := buf.(protobuf.CacheDeleteReq)
		err := h.module.Delete(req.Name , req.Key)
		resp := protobuf.CacheDeleteResp{Status:protobuf.Status{Error:err}}
		message.AddResponse(resp)
	}

	case protobuf.CacheTestConnectionReqNum:{
		//TestConnectionReq
		req := buf.(protobuf.CacheTestConnectionReq)
		err := h.module.TestConnection(req.Name)
		resp := protobuf.CacheTestConnectionResp{Status:protobuf.Status{Error:err}}
		message.AddResponse(resp)
	}

	default:{
		//Log the error
		mylog.Error("Request type error:",syserr.RequestError{})
		message.AddResponse(protobuf.ErrorResp{Status:protobuf.Status{Error:syserr.RequestError{}}})
	}
	}
}

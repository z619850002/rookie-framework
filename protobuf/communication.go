package protobuf

type ClientMessage interface {

}

type CSRequest struct {
	ClientID 	string
	MessageProtocolNum 	int
	Message 	ClientMessage
}

func (h CSRequest) GetProtocolNum()int{
	return RequestNum
}




type CSResponse struct {
	ClientID 	string
	MessageProtocolNum 	int
	Message 	ClientMessage
}



func (h CSResponse) GetProtocolNum()int{
	return ResponseNum
}
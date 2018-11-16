package protobuf

type ErrorResp struct {
	Status	Status
}

func (h ErrorResp) GetProtocolNum() int{
	return ErrorNum
}


type EmptyResp struct {
	Status Status
}

func (h EmptyResp) GetProtocolNum()	int{
	return EmptyRespNum
}
package protobuf

import "time"

//CacheGetReq is the request send to cache module to get the value by key,
//NO. 0.
type CacheGetReq struct {
	//The name of the cache source.
	Name 	string
	//The key of the key-value pair.
	Key 	string
}


//GetProtocolNum returns the protocol number of this request.
func (h CacheGetReq) GetProtocolNum() int{
	return CacheGetReqNum
}

type CacheGetResp struct {
	Value 	string
	Status	Status
}

func (h CacheGetResp) GetProtocolNum() int{
	return CacheGetRespNum
}



//Set
type CacheSetReq struct {
	//Name is the name of the cache source.
	Name 	string
	Key 	string
	Value 	string
	ExpireTime time.Duration
}

func (h CacheSetReq) GetProtocolNum() int{
	return CacheSetReqNum
}

type CacheSetResp struct {
	Status Status
}

func (h CacheSetResp) GetProtocolNum() int{
	return CacheSetRespNum
}


//Delete

type CacheDeleteReq struct {
	Name 	string
	Key 	string
}

func (h CacheDeleteReq) GetProtocolNum() int{
	return CacheDeleteReqNum
}

type CacheDeleteResp struct {
	Status 	Status
}

func (h CacheDeleteResp) GetProtocolNum() int{
	return CacheDeleteRespNum
}


//TestConnection

type CacheTestConnectionReq struct {
	Name	string
}

func (h CacheTestConnectionReq) GetProtocolNum() int{
	return CacheTestConnectionReqNum
}

type CacheTestConnectionResp struct {
	Status Status
}

func (h CacheTestConnectionResp) GetProtocolNum() int{
	return CacheTestConnectionRespNum
}
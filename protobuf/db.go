package protobuf

import (
	"database/sql"
)

//Exec
type DBExecReq struct {
	DsName string
	StrSQL string
	Args []interface{}
}

func (h DBExecReq) GetProtocolNum() int {
	return DBExecReqNum
}

type DBExecResp struct {
	Status Status
}

func (h DBExecResp) GetProtocolNum() int {
	return DBExecRespNum
}

//Query
type DBQueryReq struct {
	DsName string
	StrSQL string
	Args []interface{}
}

func (h DBQueryReq) GetProtocolNum() int {
	return DBQueryReqNum
}

type DBQueryResp struct {
	Rows *sql.Rows
	Status Status
}

func (h DBQueryResp) GetProtocolNum() int {
	return DBQueryRespNum
}

//QueryRow
type DBQueryRowReq struct {
	DsName string
	StrSQL string
	Args []interface{}
}

func (h DBQueryRowReq) GetProtocolNum() int {
	return DBQueryRowReqNum
}

type DBQueryRowResp struct {
	Row *sql.Row
	Status Status
}

func (h DBQueryRowResp) GetProtocolNum() int {
	return DBQueryRowRespNum
}
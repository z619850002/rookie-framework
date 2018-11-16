package config

import (
	"testing"
	"fmt"
)


type SysParameters struct {
	//parameters for websocket
	Port              int   // Open Port of the service.
	WriteWait         int   // Milliseconds until write times out.
	PongWait          int   // Timeout for waiting on pong.
	PingPeriod        int   // Milliseconds between pings.
	MaxMessageSize    int64 // Maximum size in bytes of a message.
	MessageBufferSize int
	////parameters for DB and Cache.
	DBInfoTable    string
	CacheInfoTable 	string
}

func (h* SysParameters) GetType() string{
	return "test"
}


func TestConfigDB_GetParameters(t *testing.T) {
	config := ConfigDB{
		DriverName: "mysql",
		Dbhost: "172.26.164.74",
		Dbport: "3306",
		Dbuser: "root",
		Dbpassword: "12345678",
		Dbname: 	"dbconfig",
		Tblname: 	"params",
	}

	param := SysParameters{}

	config.GetParameters(&param)

	fmt.Println(param)


}

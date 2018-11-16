package rookie

import (
	"testing"
	"hub000.xindong.com/rookie/rookie-framework/gate/demo"
	"hub000.xindong.com/rookie/rookie-framework/gate"
	"time"
	"hub000.xindong.com/rookie/rookie-framework/db"
	"database/sql"
	"fmt"
	"sausage-server/model"
)

func TestManager(t *testing.T) {
	p := demo.NewProcessor()
	c := gate.Config{
		Port:8080,
		WriteWait: 10,
		PongWait:  	time.Second * 60,// Timeout for waiting on pong.
		PingPeriod:        	time.Second * 54, 	// Milliseconds between pings.
		MaxMessageSize:    	512,         	// Maximum size in bytes of a message.
		MessageBufferSize: 	256,           	// The max amount of messages that can be in a WsConn buffer before it starts dropping them.
	}

	m := NewServer(p , &c).UseCache().UseDataBase()

	dburl := "root" + ":" + "12345678" + "@tcp(" + "172.26.164.74" + ":" + "3306" + ")/" + "sausage_db" + "?charset=utf8"
	localDB, err := sql.Open("mysql", dburl)
	if err != nil {
		fmt.Println(err)
	}

	db.Module.RegisterDataSource("base" , db.NewSQLDataSource(localDB))

	res , err := db.Module.Query("base" , "SELECT * FROM gun")
	if (err != nil){
		t.Error(err)
	}
	for res.Next(){
		gun := model.Gun{}
		res.Scan(&gun.ID , &gun.Name , &gun.Firepower , &gun.ReloadSpeed , &gun.Stability , &gun.Recoil , &gun.Damage  , &gun.Price)
		fmt.Println(gun)
	}


	m.Run()
}

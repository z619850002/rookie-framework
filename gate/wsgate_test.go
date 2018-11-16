package gate

import (
	"encoding/binary"
	"flag"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"hub000.xindong.com/rookie/rookie-framework/protobuf"
	"log"
	"net/http"
	"net/url"
	"sausage-shoot-proto/protocol"
	"testing"
)

type TestHandler struct {}

func (m *TestHandler)Unmarshal(t int, message []byte, clientID string) (protobuf.CSRequest, error){
	fmt.Println("in Unmarshal func")
	return protobuf.CSRequest{}, errors.New("sdf")
}

func (m *TestHandler)Marshal(resp protobuf.CSResponse) (*Message, error){
	return &Message{}, nil
}


func TestWSGate(t *testing.T) {
	sigChan := make(chan string, 1)

	go func() {
		c := &Config{
			Port:10000,
			WriteWait:20,
			PongWait:60,
			PingPeriod:54,
			MaxMessageSize:512,
			MessageBufferSize:256,
		}
		g := &TestHandler{}
		gate := NewGate(c).UseWS(g).UseHttp()
		gate.RegisterHttpRouter("/index", func(writer http.ResponseWriter, r *http.Request) {
			fmt.Println("this is a http request")
			sigChan <- "over"
		})
		gate.StartServer()
	}()

	GetWSConn()

	Get("http://localhost:10000/index")
	fmt.Println(<- sigChan)
}


func GetWSConn(){
	var addr = flag.String("addr", "localhost:10000", "http service address")

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	c.WriteMessage(websocket.BinaryMessage, createMsg())
}

func createMsg() []byte {
	// build modulemsg
	login := protocol.GetConfReq{
		PlayerID: "qinhan",
	}
	data, err := proto.Marshal(&login)
	if err != nil {
		log.Println(err)
	}

	// add protocol num
	m := make([]byte, 2+len(data))
	// 默认使用大端序
	binary.BigEndian.PutUint16(m, uint16(4))
	copy(m[2:], data)

	return m
}

func getMessage(msg []byte) (uint16, []byte) {
	temp := make([]byte, 2)
	body := make([]byte, len(msg)-2)
	copy(temp, msg[:2])
	copy(body, msg[2:])
	top := binary.BigEndian.Uint16(temp)
	return top, body
}
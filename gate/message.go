package gate

type Message struct {
	T   int // websocket message type : websocket.TextMessage / websocket.BinaryMessage
	Msg []byte
}

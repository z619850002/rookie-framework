package gate

import (
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"hub000.xindong.com/rookie/rookie-framework/log"
	"hub000.xindong.com/rookie/rookie-framework/syserr"
)

// WsConn wrapper around websocket connections.
type WsConn struct {
	Request *http.Request
	conn    *websocket.Conn
	output  chan *Message
	closeSig chan *Message
	wsgate    *WSGate
	isRunning    bool
	rwMutex *sync.RWMutex
	player  string
}

func NewWsConn(r *http.Request, conn *websocket.Conn, g *WSGate) *WsConn {
	return &WsConn{
		Request: r,
		conn:    conn,
		output:  make(chan *Message, g.gate.Config.MessageBufferSize),
		closeSig: make(chan *Message),
		wsgate:    g,
		isRunning:    true,
		rwMutex: &sync.RWMutex{},
	}
}

func (w *WsConn) SetPlayer (p string) {
	w.player = p
}

// WriteMessage :
func (w *WsConn) WriteMessage(message *Message) {
	if !w.IsRunning() {
		mylog.Error("WriteMessage fail.", syserr.WriteFailError{})
		return
	}

	w.output <- message
}

func (w *WsConn) run() {
	//register in hub
	//TODO : wait until get the response of the hub
	w.wsgate.hub.register <- w

	go w.writePump()
	w.readPump()

	//execute close func
	if w.IsRunning() {
		w.close()
	}
}

func (w *WsConn) writePump() {
	ticker := time.NewTicker(time.Duration(w.wsgate.gate.Config.PingPeriod) * time.Second)
	defer ticker.Stop()

loop:
	for {
		select {
		case <-ticker.C:
			w.ping()
		case msg, ok := <-w.closeSig:
			if ok {
				w.writeRaw(msg)
			}
			break loop

		case msg, ok := <-w.output:
			if !ok {
				break loop
			}

			err := w.writeRaw(msg)

			if err != nil {
				mylog.Error("writeRaw fail.", err)
				break loop
			}
		}
	}
}

func (w *WsConn) readPump() {
	w.conn.SetReadLimit(w.wsgate.gate.Config.MaxMessageSize)
	w.conn.SetReadDeadline(time.Now().Add(time.Duration(w.wsgate.gate.Config.PongWait) * time.Second))

	w.conn.SetPongHandler(func(string) error {
		return w.conn.SetReadDeadline(time.Now().Add(time.Duration(w.wsgate.gate.Config.PongWait) * time.Second))
	})

	for {
		t, message, err := w.conn.ReadMessage()

		if err != nil {
			mylog.Error("Websocket ReadMessage fail.", err)
			break
		}

		w.wsgate.processor.MessageHandler(w, t , message)
	}
}

func (w *WsConn) ping() {
	w.writeRaw(&Message{T: websocket.PingMessage, Msg: []byte{}})
}

// IsClosed returns the status of the connection.
func (w *WsConn) IsRunning() bool {
	w.rwMutex.Lock()
	defer w.rwMutex.Unlock()

	return w.isRunning
}

func (w *WsConn) writeRaw(message *Message) error {
	if !w.IsRunning() {
		return syserr.SendOnClosedWsConnError{}
	}

	w.conn.SetWriteDeadline(time.Now().Add(time.Duration(w.wsgate.gate.Config.WriteWait) * time.Second))
	err := w.conn.WriteMessage(message.T, message.Msg)

	// execute close func
	if message.T == websocket.CloseMessage || err != nil {
		w.close()
	}

	return err
}

func (w *WsConn) Close() {
	if w.isRunning {
		w.close()
	}
}

// CloseWithMsg closes the wsConn with the provided payload.
// Use the FormatCloseMessage function to format a proper close message payload.
func (w *WsConn) CloseWithMsg(msg []byte) {
	if w.IsRunning() {
		w.closeSig <- &Message{T: websocket.CloseMessage, Msg: msg}
		w.close()
	}
}

//send all the
func (w *WsConn) close() {
	if w.IsRunning() {
		w.rwMutex.Lock()
		w.isRunning = false
		w.rwMutex.Unlock()

		w.conn.Close()
		close(w.output)
		//remove the connection out of the hub pool.
		if w.wsgate.hub.IsRunning() {
			//TODO : wait until get the response of the hub
			w.wsgate.hub.unregister <- w
		}
	}
}

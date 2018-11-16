package gate

import (
	"github.com/gorilla/websocket"
	"hub000.xindong.com/rookie/rookie-framework/log"
	"hub000.xindong.com/rookie/rookie-framework/syserr"
	"net/http"
)

// Gate implements a websocket manager.
type WSGate struct {
	gate 		*Gate
	Upgrader  	*websocket.Upgrader
	processor 	*Processor
	hub       	*hub
}

// New creates a new Gate instance with default Upgrader and Config.
func NewWSGate(h MsgHandler) *WSGate {
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	hub := newHub()

	go hub.run()

	gate := &WSGate{
		Upgrader:  	upgrader,
		processor: 	NewProcessor(h),
		hub:       	hub,
	}

	gate.processor.g = gate

	return gate
}

func (g *WSGate) StartService() {
	g.gate.srvMux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		mylog.Info("websocket get connected." + r.RemoteAddr)
		g.handleRequest(w, r)
	})
}

// HandleRequest upgrades http requests to websocket connections and dispatches them to be handled by the Gate instance.
func (g *WSGate) handleRequest(w http.ResponseWriter, r *http.Request) {
	if !g.hub.IsRunning() {
		mylog.Fatal("gate instance is closed")
		return
	}

	conn, err := g.Upgrader.Upgrade(w, r, w.Header())

	if err != nil {
		mylog.Info("upgrade fail.", err)
		return
	}

	wsClient := NewWsConn(r, conn, g)

	//TODO : set id
	wsClient.player = wsClient.Request.RemoteAddr

	wsClient.run()
}

// Close closes the Gate instance and all connected wsConn.
func (g *WSGate) Close() error {
	return g.CloseWithMsg([]byte{})
}

// CloseWithMsg closes the Gate instance with the given close payload and all connected WsConn.
// Use the FormatCloseMessage function to format a proper close message payload.
func (g *WSGate) CloseWithMsg(msg []byte) error {
	if !g.hub.IsRunning() {
		return syserr.ClosedGate{}
	}

	g.hub.Close()

	return nil
}

//SendToPlayer : send to player by unique ID
func (g *WSGate) SendToPlayer(uniqueID string, message *Message) error {
	w, err := g.hub.GetPlayerByID(uniqueID)
	if err != nil {
		return err
	}
	w.WriteMessage(message)
	return nil
}

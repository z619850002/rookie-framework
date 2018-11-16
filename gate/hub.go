package gate

import (
	"sync"
	"hub000.xindong.com/rookie/rookie-framework/syserr"
)

type hub struct {
	connections map[*WsConn]bool
	register    chan *WsConn
	unregister  chan *WsConn
	exit        chan struct{}
	isRunning    	bool
	rwMutex     *sync.RWMutex
}

func newHub() *hub {
	return &hub{
		connections: make(map[*WsConn]bool),
		register:    make(chan *WsConn),
		unregister:  make(chan *WsConn),
		exit:        make(chan struct{}),
		isRunning:    	 true,
		rwMutex:     &sync.RWMutex{},
	}
}

func (h *hub) run() {
loop:
	for {
		select {
		case <-h.exit:
			break loop
		case s := <-h.register:
			h.rwMutex.Lock()
			h.connections[s] = true
			h.rwMutex.Unlock()
		case s := <-h.unregister:
			if _, ok := h.connections[s]; ok {
				h.rwMutex.Lock()
				delete(h.connections, s)
				h.rwMutex.Unlock()
			}
		}
	}
}

func (h *hub) Close() {
	h.rwMutex.Lock()
	defer h.rwMutex.Unlock()

	for conn:= range h.connections {
		conn.Close()
	}
	h.isRunning = false
	h.exit <- struct{}{}
}

func (h *hub) IsRunning() bool {
	h.rwMutex.RLock()
	defer h.rwMutex.RUnlock()
	return h.isRunning
}

func (h *hub) len() int {
	h.rwMutex.RLock()
	defer h.rwMutex.RUnlock()

	return len(h.connections)
}

func (h *hub) GetPlayerByID (ID string) (*WsConn, error) {
	for conn := range h.connections {
		if  conn.player == ID {
			if conn.IsRunning() {
				return conn,nil
			}
			return nil, syserr.ClosedWsConnError{}
		}
	}
	return nil, syserr.NoMatchedPlayerError{}
}




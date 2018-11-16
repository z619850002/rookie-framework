package gate

import (
	"errors"
	"net/http"
	"sync"
)

type HttpGate struct {
	gate 		*Gate
	r			map[string]func(http.ResponseWriter, *http.Request)
	sync.Mutex
}

func NewHttpGate() *HttpGate {
	return &HttpGate{
		r : make(map[string]func(http.ResponseWriter, *http.Request)),
	}
}

func (h *HttpGate) SetGate(g *Gate) {
	h.gate = g
}

func (h *HttpGate) StartHttpServer (){
	for r, f := range h.r {
		h.gate.srvMux.HandleFunc(r, f)
	}
}

func (h *HttpGate) Register(r string, f func(http.ResponseWriter, *http.Request)) error {
	h.Lock()
	defer h.Unlock()
	if _, ok := h.r[r]; ok {
		return errors.New("registered http route")
	} else {
		h.r[r] = f
		return nil
	}
}
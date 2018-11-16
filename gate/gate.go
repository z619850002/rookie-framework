package gate

import (
	"fmt"
	"github.com/pkg/errors"
	"hub000.xindong.com/rookie/rookie-framework/log"
	"net/http"
)

type Gate struct {
	Config 		*Config
	http 		*HttpGate
	ws 			*WSGate
	srvMux		*http.ServeMux
}

func NewGate(c *Config) *Gate {
	return &Gate{
		Config:c,
		srvMux: http.NewServeMux(),
		http:nil,
		ws:nil,
	}
}

func (g *Gate) UseWS(h MsgHandler) *Gate{
	ws := NewWSGate(h)
	ws.gate = g
	g.ws = ws
	return g
}

func (g *Gate) UseHttp() *Gate{
	http := NewHttpGate()
	http.SetGate(g)
	g.http = http
	return g
}

func (g *Gate) StartServer(){
	if g.ws != nil || g.http != nil {
		addr := ":" + fmt.Sprintf("%d", g.Config.Port)
		if g.ws != nil {
			g.ws.StartService()
		}
		if g.http != nil {
			g.http.StartHttpServer()
		}
		//http.ListenAndServe(addr, g.srvMux)
		if err := http.ListenAndServe(addr, g.srvMux); err != nil {
			mylog.Fatal("start server fail.", err)
		}
		mylog.Info("start server success.")
	} else {
		mylog.Fatal("no service.")
	}
}

func (g *Gate) RegisterHttpRouter(r string, f func(http.ResponseWriter, *http.Request)){
	if g.http != nil {
		err := g.http.Register(r,f)
		if err != nil {
			mylog.Fatal("RegisterHttpRouter Fail.", err)
		}
	} else {
		mylog.Fatal("http server locked!")
	}
}

func (g *Gate) GetWsRouter() (*Router, error) {
	if g.ws != nil {
		return g.ws.processor.Router, nil
	}
	return nil, errors.New("ws server locked!")
}
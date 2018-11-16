package rookie

import (
	"hub000.xindong.com/rookie/rookie-framework/cache"
	"hub000.xindong.com/rookie/rookie-framework/db"
	"hub000.xindong.com/rookie/rookie-framework/gate"
	)

type Server struct {
	g           *gate.Gate
}

func NewServer(gate *gate.Gate) *Server {
	return &Server{g: gate}
}

//UseCache will init the cache package
func (h *Server) UseCache() *Server {
	//Initialize the cache.
	cache.OnInit()
	return h
}

//UseCache will init the db package
func (h *Server) UseDataBase() *Server {
	//initialize the db.
	db.OnInit()
	return h
}

//Run will run all modules in the modulesPool
func (h *Server) Run() {

	h.g.StartServer()
}

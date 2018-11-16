package module

import (
	"hub000.xindong.com/rookie/rookie-framework/skeleton"
	"hub000.xindong.com/rookie/rookie-framework/communication"
	"sync"
)

//Module is the basic component of code.
type Module struct {
	skeleton.Skeleton
	communication.Communicator
	handler *Handler
	//If this module is running
	running bool
	closeChan	chan bool
}


func NewModule() *Module{
	return &Module{
		Skeleton:skeleton.NewBaseSkeleton(),
		running:false,
		//The channel to kill the go routine in this module.
		closeChan: make(chan bool),
	}
}

//SetCommunicator takes in the pointer of the communicator and bind the communicator
// with this module.
func (h *Module) SetCommunicator(comm communication.Communicator){
	h.Communicator = comm
	//Until this module is running this communicator can unlock.
	comm.Lock()
}

func (h *Module) SetHandler(handler Handler){
	//Bind handler and module together.
	h.handler = &handler
}

//Run will run the main thread in this module.
func (h *Module) Run(){
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		h.running = true
		//Unlock the communicator.
		h.Communicator.UnLock()
		wg.Done()
		//The main thread in this module.
		Loop:
		for ;;{
			select {
			//Shut down this module
			case sig :=  <- h.closeChan:{
				//Jump out the loop.
				if (sig){
					break Loop
				}
			}
			//The GetInput is derived from the rpc communicator.
			case m,ok := <- h.Communicator.GetInputChan():{
				if (ok){
					//Run in a new go routine.
					go func (){
						//Handler the message
						(*h.handler).Handle(m)
					}()
				}
			}
			}
		}
		//TODO:close elegantly.
		//Lock the communicator.
		oldChan := h.GetInputChan()
		h.Communicator.Lock()
		h.running = false
		//Clear the buffer in the input channel.
		Loop2:
		for {
			select {
				case m,ok := <- oldChan:{
					if (!ok){
						break Loop2
					}
					(*h.handler).Handle(m)
				}
			default:
				break
			}
		}
	}()
	wg.Wait()
}

//If this module is running.
func (h *Module) IsRunning()bool{
	return h.running
}

//Close will shut down this module.
func (h *Module) Close(){
	h.closeChan <- true
	//Wait until this module closed.
	for h.running{

	}
}
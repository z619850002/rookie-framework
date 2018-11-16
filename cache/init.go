package cache

import (
			"hub000.xindong.com/rookie/rookie-framework/communication"
	)

var Module *CacheModule
var Communicator communication.Communicator

func OnInit(){
	Module = NewCacheModule()
	Communicator = communication.NewRPCCommunicator()
	Module.SetCommunicator(Communicator)
	Module.SetHandler(&Handler{module:Module})
	Module.Run()
}

package db

import (
			"hub000.xindong.com/rookie/rookie-framework/communication"
	)

var Module *DBModule
var Communicator communication.Communicator

func OnInit() {
	Module = NewDBModule()
	Communicator = communication.NewRPCCommunicator()
	Module.SetCommunicator(Communicator)
	Module.SetHandler(&Handler{module: Module})
	Module.Run()
}

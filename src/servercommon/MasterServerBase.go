package servercommon

import (
	"./syshandler"

	"../base/ginterface"
	"../base/server"
)

// MasterServerBase implements the common operations of the master server
type MasterServerBase struct {
	*server.GameServer // base class
}

// OnRegisterToMaster is called for registering to the master server
func (s *MasterServerBase) OnRegisterToMaster() {

}

// NewMasterServerBase is a constructor of MasterServerBase
func NewMasterServerBase(hook ginterface.IGameServerHook, serverType int, port int, serverName string, log ginterface.IGameLogger) *MasterServerBase {
	ret := &MasterServerBase{}
	ret.GameServer = server.NewGameServer(hook, serverType, port, serverName, log)
	ret.RegisterHandler(syshandler.NewRegisterSubServerHandler(ret))
	return ret
}

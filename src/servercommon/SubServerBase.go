package servercommon

import (
	"../base/ginterface"
	"../base/server"
	"./sysdefine"
	"./syshandler"
)

// SubServerBase implements the common operations of the subserver
type SubServerBase struct {
	*server.GameServer // base class
}

// OnRegisterToMaster is called for registering to the master server
func (s *SubServerBase) OnRegisterToMaster() {
	packet := sysdefine.NewRegisterSubServerPacket()
	packet.ServerType = s.Setting.ServerType
	s.SendPacket(s.MasterPeer, packet)
}

// NewSubServerBase is a constructor of SubServerBase
func NewSubServerBase(hook ginterface.IGameServerHook, serverType int, port int, serverName string, log ginterface.IGameLogger) *SubServerBase {
	ret := &SubServerBase{}
	ret.GameServer = server.NewGameServer(hook, serverType, port, serverName, log)
	ret.RegisterHandler(syshandler.NewRegisterSubServerResultHandler(ret))
	return ret
}

package servercommon

import (
	"../base/ginterface"
	"../base/server"
	"./sysdefine"
	"./syshandler"
	"./sysinfo"
	uuid "github.com/satori/go.uuid"
)

// SubServerBase implements the common operations of the subserver
type SubServerBase struct {
	*server.GameServer // base class

	Players map[uuid.UUID]sysinfo.PlayerInfo // map of players with key denoting the player key
}

// GetModule returns the specific module to resolve import cycle
func (s *SubServerBase) GetModule(m interface{}) interface{} {
	ret := s.GameServer.GetModule(m)
	if ret != nil {
		return ret
	}
	switch m.(type) {
	case map[uuid.UUID]sysinfo.PlayerInfo:
		return s.Players
	}
	return nil
}

// OnRegisterToMaster is called for registering to the master server
func (s *SubServerBase) OnRegisterToMaster() {
	packet := sysdefine.NewRegisterSubServerPacket()
	packet.ServerType = s.Setting.ServerType
	packet.Port = s.Setting.Port
	packet.ServerName = s.Setting.ServerName
	s.SendPacket(s.MasterPeer, packet)
}

// NewSubServerBase is a constructor of SubServerBase
func NewSubServerBase(hook ginterface.IGameServerHook, serverType int, port int, serverName string, log ginterface.IGameLogger) *SubServerBase {
	ret := &SubServerBase{}
	ret.GameServer = server.NewGameServer(hook, serverType, port, serverName, log)
	ret.Players = make(map[uuid.UUID]sysinfo.PlayerInfo)
	ret.RegisterHandler(syshandler.NewRegisterSubServerResultHandler(ret))
	return ret
}

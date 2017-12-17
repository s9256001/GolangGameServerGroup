package servercommon

import (
	"./syshandler"
	uuid "github.com/satori/go.uuid"

	"../base/ginterface"
	"../base/server"
	"./sysinfo"
)

// MasterServerBase implements the common operations of the master server
type MasterServerBase struct {
	*server.GameServer // base class

	SubServers map[uuid.UUID]sysinfo.SubServerInfo // map of subservers with key denoting the peer id
	Players    map[uuid.UUID]sysinfo.PlayerInfo    // map of players with key denoting the player key
}

// GetModule returns the specific module to resolve import cycle
func (s *MasterServerBase) GetModule(m interface{}) interface{} {
	ret := s.GameServer.GetModule(m)
	if ret != nil {
		return ret
	}
	switch m.(type) {
	case map[uuid.UUID]sysinfo.SubServerInfo:
		return s.SubServers
	case map[uuid.UUID]sysinfo.PlayerInfo:
		return s.Players
	}
	return nil
}

// OnRegisterToMaster is called for registering to the master server
func (s *MasterServerBase) OnRegisterToMaster() {

}

// NewMasterServerBase is a constructor of MasterServerBase
func NewMasterServerBase(hook ginterface.IGameServerHook, serverType int, port int, serverName string, log ginterface.IGameLogger) *MasterServerBase {
	ret := &MasterServerBase{}
	ret.GameServer = server.NewGameServer(hook, serverType, port, serverName, log)
	ret.SubServers = make(map[uuid.UUID]sysinfo.SubServerInfo)
	ret.Players = make(map[uuid.UUID]sysinfo.PlayerInfo)
	ret.RegisterHandler(syshandler.NewRegisterSubServerHandler(ret))
	return ret
}

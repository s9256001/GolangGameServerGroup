package server

import (
	"golang.org/x/net/websocket"

	"../../base/ginterface"
	//"../../gamecommon/gamedefine"
	"../../servercommon"
	"../../servercommon/sysdefine"
	"../handler"
	"../peer"
)

// DerivedGameServer implements the cache server instance
type DerivedGameServer struct {
	*servercommon.MasterServerBase // base class
}

// OnStart is called before Start()
func (s *DerivedGameServer) OnStart() {
	// TODO: test
	// for output serialized string
	// testPacket := gamedefine.NewRegisterAccountPacket()
	// testPacket.Account = "s9256001"
	// testPacket.Password = "s9256001"
	// str, _ := json.Marshal(testPacket)
	// s.GetLog().Debug("%s\n", str)

	s.GetLogger().Debug("DerivedGameServer.OnStart: serverName = %s\n", s.Setting.ServerName)
}

// OnStopped is called at the end of Stop()
func (s *DerivedGameServer) OnStopped() {
	s.GetLogger().Debug("DerivedGameServer.OnStopped: serverName = %s\n", s.Setting.ServerName)
}

// OnCreatePeer is called to create the custom peer
func (s *DerivedGameServer) OnCreatePeer(conn *websocket.Conn) ginterface.IGamePeer {
	return peer.NewDerivedGamePeer(s, conn)
}

// OnDefaultHandle is called when there is no corresponding packet handler
func (s *DerivedGameServer) OnDefaultHandle(peer ginterface.IGamePeer, info string) {
	s.GetLogger().Debug("DerivedGameServer.OnDefaultHandle: peerID = %s, info = %s\n", peer.GetPeerID().String(), info)
}

// NewDerivedGameServer is a constructor of DerivedGameServer
func NewDerivedGameServer() *DerivedGameServer {
	log := servercommon.NewConsoleGameLogger()
	ret := &DerivedGameServer{}
	// todo
	ret.MasterServerBase = servercommon.NewMasterServerBase(ret, sysdefine.ServerTypeCache, 7770, "cache", log)
	ret.RegisterHandler(handler.NewRegisterAccountHandler(ret))
	ret.RegisterHandler(handler.NewLoginHandler(ret))
	ret.RegisterHandler(handler.NewEnterRegionHandler(ret))
	return ret
}

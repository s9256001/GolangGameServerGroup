package server

import (
	"golang.org/x/net/websocket"

	"../../base/ginterface"
	"../../servercommon"
	"../../servercommon/sysdefine"
	testgame "../../testgame/game"
	"../handler"
	"../peer"
)

// DerivedGameServer implements the region server instance
type DerivedGameServer struct {
	*servercommon.SubServerBase // base class

	Games map[int]ginterface.IGame // map of games with key denoting the gameid
}

// GetModule returns the specific module to resolve import cycle
func (s *DerivedGameServer) GetModule(m interface{}) interface{} {
	ret := s.SubServerBase.GetModule(m)
	if ret != nil {
		return ret
	}
	switch m.(type) {
	case map[int]ginterface.IGame:
		return s.Games
	}
	return nil
}

// OnStart is called before Start()
func (s *DerivedGameServer) OnStart() {
	s.GetLogger().Debug("OnStart: serverName = %s\n", s.Setting.ServerName)

	var game ginterface.IGame
	game = testgame.NewDerivedGame(s.GetLogger(), s)
	s.Games[game.GameID()] = game

	for _, game := range s.Games {
		if game.Init(nil) == true {
			s.GetLogger().Debug("OnStart: game %d initialize success.\n", game.GameID())
		} else {
			s.GetLogger().Error("OnStart: game %d initialize failed.\n", game.GameID())
		}
	}
}

// OnStopped is called at the end of Stop()
func (s *DerivedGameServer) OnStopped() {
	s.GetLogger().Debug("OnStopped: serverName = %s\n", s.Setting.ServerName)
}

// OnCreatePeer is called to create the custom peer
func (s *DerivedGameServer) OnCreatePeer(conn *websocket.Conn) ginterface.IGamePeer {
	return peer.NewDerivedGamePeer(s, conn)
}

// OnDefaultHandle is called when there is no corresponding packet handler
func (s *DerivedGameServer) OnDefaultHandle(peer ginterface.IGamePeer, info string) {
	s.GetLogger().Debug("OnDefaultHandle: info = %s\n", info)
}

// NewDerivedGameServer is a constructor of DerivedGameServer
func NewDerivedGameServer() *DerivedGameServer {
	log := servercommon.NewConsoleGameLogger()
	ret := &DerivedGameServer{
		Games: make(map[int]ginterface.IGame),
	}
	ret.SubServerBase = servercommon.NewSubServerBase(ret, sysdefine.ServerTypeRegion, 7772, "region", log)
	ret.Setting.MasterURL = "ws://127.0.0.1:7770/cache"
	ret.RegisterHandler(handler.NewEnterRegionHandler(ret))
	ret.RegisterHandler(handler.NewEnterRegionResultHandler(ret))
	ret.RegisterHandler(handler.NewEnterGameHandler(ret))
	return ret
}

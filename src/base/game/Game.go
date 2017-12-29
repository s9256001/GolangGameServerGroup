package game

import (
	"encoding/json"

	"../basedefine"
	"../ginterface"
	"../node"
)

// Game is an abstract class of the game
// It fits the interface of IGame, and depends on the derived type to implement the interface of IGameHook
type Game struct {
	ginterface.IGameHook // hook
	*node.Node           // base class

	Handlers map[int]ginterface.IGameHandler // map of packet handlers with key denoting the packet command

	// todo
	Server ginterface.IGameServer // game server
}

// GameID gets the game ID of the game
func (g *Game) GameID() int {
	return 0
}

// RegisterHandler registers the packet handler
func (g *Game) RegisterHandler(handler ginterface.IGameHandler) {
	g.Handlers[handler.Code()] = handler
}

// Init initlizes the game
func (g *Game) Init(setting interface{}) bool {
	g.Log.Debug("Init: initialization start\n")
	return g.OnInit(setting)
}

// Release releases the game
func (g *Game) Release() bool {
	g.Log.Debug("Release: release\n")
	return g.OnRelease()
}

// HandlePacket handles the packet
func (g *Game) HandlePacket(peer ginterface.IGamePeer, info string) {
	// deserialize the packet
	basePacket := basedefine.GameBasePacket{}
	if err := json.Unmarshal([]byte(info), &basePacket); err != nil {
		g.Log.Error("Game.HandlePacket: deserialized failed! err = %v\n", err)
		return
	}
	// dispatch the packet to handlers
	if handler, ok := g.Handlers[basePacket.Code]; ok {
		g.Log.Trace("Game.GameHandler.Handle: code = %d, info = %s\n", handler.Code(), info)
		handler.Handle(peer, info)
	} else {
		g.OnDefaultHandle(peer, info)
	}
}

// GetServer returns the game server
func (g *Game) GetServer() ginterface.IGameServer {
	return g.Server
}

// NewGame is a constructor of Game
func NewGame(hook ginterface.IGameHook, log ginterface.IGameLogger, server ginterface.IGameServer) *Game {
	ret := &Game{
		IGameHook: hook,
		Node:      node.NewNode(log),

		Handlers: make(map[int]ginterface.IGameHandler),

		Server: server,
	}
	return ret
}

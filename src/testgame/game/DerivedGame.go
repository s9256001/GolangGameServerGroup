package game

import (
	"../../base/game"
	"../../base/ginterface"
	"../../gamecommon/gamedefine"
	"../handler"
	"../model"
	"../state"
)

// DerivedGame implements the test game instance
type DerivedGame struct {
	*game.Game // base class
}

// GameID gets the game ID of the game
func (g *DerivedGame) GameID() int {
	return gamedefine.GameIDTestGame
}

// OnInit is called when Init()
func (g *DerivedGame) OnInit(setting interface{}) bool {
	return true
}

// OnRelease is called when Release()
func (g *DerivedGame) OnRelease() bool {
	return true
}

// OnDefaultHandle is called when there is no corresponding packet handler
func (g *DerivedGame) OnDefaultHandle(peer ginterface.IGamePeer, info string) {
	g.Log.Debug("DerivedGame.OnDefaultHandle: info = %s\n", info)
}

// NewDerivedGame is a constructor of DerivedGame
func NewDerivedGame(log ginterface.IGameLogger, server ginterface.IGameServer) *DerivedGame {
	ret := &DerivedGame{}
	ret.Game = game.NewGame(ret, log, server)
	ret.Model = model.NewDerivedModel()

	ret.State = state.NewInitState(ret)
	ret.RegisterHandler(handler.NewEnterGameHandler(ret))
	return ret
}

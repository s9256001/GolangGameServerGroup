package handler

import (
	"encoding/json"

	"../../base/ginterface"
	"../../base/handler"
	"../../gamecommon/gamedefine"
	"../../servercommon/sysdefine"
	"../../servercommon/sysinfo"
	uuid "github.com/satori/go.uuid"
)

// EnterRegionHandler handles player's entering region
type EnterRegionHandler struct {
	*handler.GameHandler // base class
}

// Code is the associated packet command
func (h *EnterRegionHandler) Code() int {
	return gamedefine.EnterRegion
}

// OnHandle is called when Handle()
func (h *EnterRegionHandler) OnHandle(peer ginterface.IGamePeer, info string) bool {
	log := h.Server.GetLogger()
	players := h.Server.GetModule(map[uuid.UUID]sysinfo.PlayerInfo{}).(map[uuid.UUID]sysinfo.PlayerInfo)

	response := gamedefine.NewEnterRegionResultPacket()

	packet := &gamedefine.EnterRegionPacket{}
	if err := json.Unmarshal([]byte(info), &packet); err != nil {
		log.Error("EnterRegionHandler.OnHandle(): failed to deserialize! info = %s\n", info)
		return false
	}

	playerKey, _ := uuid.FromString(packet.PlayerKey)
	playerInfo := players[playerKey].PlayerInfoBase

	response.PeerID = packet.PeerID
	response.PlayerInfo = playerInfo
	response.Result = sysdefine.OK
	h.Server.SendPacket(peer, response)
	return true
}

// NewEnterRegionHandler is a constructor of EnterRegionHandler
func NewEnterRegionHandler(server ginterface.IGameServer) *EnterRegionHandler {
	ret := &EnterRegionHandler{}
	ret.GameHandler = handler.NewGameHandler(ret, server)
	return ret
}

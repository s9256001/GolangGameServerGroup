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

// EnterRegionHandler handles the request of player's entering region
type EnterRegionHandler struct {
	*handler.GameHandler // base class
}

// Code is the associated packet command
func (h *EnterRegionHandler) Code() int {
	return gamedefine.EnterRegion
}

// OnHandle is called when Handling the packet
func (h *EnterRegionHandler) OnHandle(peer ginterface.IGamePeer, info string) bool {
	log := h.Node.(ginterface.IGameServer).GetLogger()
	players := h.Node.(ginterface.IGameServer).GetModule(map[uuid.UUID]*sysinfo.PlayerInfo{}).(map[uuid.UUID]*sysinfo.PlayerInfo)

	response := gamedefine.NewEnterRegionResultPacket()
	response.Result = sysdefine.Failed

	defer h.Node.(ginterface.IGameServer).SendPacket(peer, response)

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
	return true
}

// NewEnterRegionHandler is a constructor of EnterRegionHandler
func NewEnterRegionHandler(node ginterface.INode) *EnterRegionHandler {
	ret := &EnterRegionHandler{}
	ret.GameHandler = handler.NewGameHandler(ret, node)
	return ret
}

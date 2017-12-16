package gamedefine

import (
	"../../base/server"
)

// RegisterAccountPacket is the packet of request of registration of the player's account
type RegisterAccountPacket struct {
	server.GameBasePacket        // base packet
	PeerID                string // player's peer id
	Account               string // player's account
	Password              string // player's password
}

// NewRegisterAccountPacket is a constructor of RegisterAccountPacket
func NewRegisterAccountPacket() *RegisterAccountPacket {
	ret := &RegisterAccountPacket{}
	ret.Code = RegisterAccount
	return ret
}

// RegisterAccountResultPacket is the packet of response of registration of player's login
type RegisterAccountResultPacket struct {
	server.GameBasePacket        // base packet
	PeerID                string `json:",omitempty"` // player's peer id
	Result                int    // result
}

// NewRegisterAccountResultPacket is a constructor of RegisterAccountResultPacket
func NewRegisterAccountResultPacket() *RegisterAccountResultPacket {
	ret := &RegisterAccountResultPacket{}
	ret.Code = RegisterAccountResult
	return ret
}

package gamedefine

import (
	"../../base/server"
	"../../servercommon/sysinfo"
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

// LoginPacket is the packet of request of player's login
type LoginPacket struct {
	server.GameBasePacket        // base packet
	PeerID                string // player's peer id
	Account               string // player's account
	Password              string // player's password
}

// NewLoginPacket is a constructor of LoginPacket
func NewLoginPacket() *LoginPacket {
	ret := &LoginPacket{}
	ret.Code = Login
	return ret
}

// LoginResultPacket is the packet of response of player's login
type LoginResultPacket struct {
	server.GameBasePacket        // base packet
	PeerID                string `json:",omitempty"` // player's peer id
	RegionAddress         string // region server's address
	PlayerKey             string // player's key for communicating to the region server
	Result                int    // result
}

// NewLoginResultPacket is a constructor of LoginResultPacket
func NewLoginResultPacket() *LoginResultPacket {
	ret := &LoginResultPacket{}
	ret.Code = LoginResult
	return ret
}

// EnterRegionPacket is the packet of request of player's entering region
type EnterRegionPacket struct {
	server.GameBasePacket        // base packet
	PeerID                string // player's peer id
	PlayerKey             string // player's key for communicating to the region server
}

// NewEnterRegionPacket is a constructor of EnterRegionPacket
func NewEnterRegionPacket() *EnterRegionPacket {
	ret := &EnterRegionPacket{}
	ret.Code = EnterRegion
	return ret
}

// EnterRegionResultPacket is the packet of response of player's entering region
type EnterRegionResultPacket struct {
	server.GameBasePacket                        // base packet
	PeerID                string                 `json:",omitempty"` // player's peer id
	PlayerInfo            sysinfo.PlayerInfoBase // the information of the player
	Result                int                    // result
}

// NewEnterRegionResultPacket is a constructor of EnterRegionResultPacket
func NewEnterRegionResultPacket() *EnterRegionResultPacket {
	ret := &EnterRegionResultPacket{}
	ret.Code = EnterRegionResult
	return ret
}

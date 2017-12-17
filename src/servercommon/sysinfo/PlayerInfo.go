package sysinfo

import (
	"github.com/satori/go.uuid"
)

// PlayerInfoBase is the basic information of the player
type PlayerInfoBase struct {
	PlayerKey uuid.UUID // player's key for communicating to the region server
	Name      string    // player's name
	Gold      int       // player's money
}

// PlayerInfo is the information of the player
type PlayerInfo struct {
	PlayerInfoBase // base class

	Data interface{} // custom data
}

package gamedefine

const (
	// RegisterAccount is the packet command of request of registration of the player's account
	RegisterAccount = iota + 101
	// RegisterAccountResult is the packet command of response of registration of the player's account
	RegisterAccountResult
	// Login is the packet command of request of player's login
	Login
	// LoginResult is the packet command of response of player's login
	LoginResult
	// Logout is the packet command of request of player's logout
	Logout
	// EnterRegion is the packet command of request of player's entering region
	EnterRegion
	// EnterRegionResult is the packet command of response of player's entering region
	EnterRegionResult
)

package ginterface

// INodeState is the state of the node
type INodeState interface {
	// OnEnter is called when entering this state
	OnEnter()
	// OnExit is called when exiting this state
	OnExit()
	// OnUpdate is called when updating this state
	OnUpdate(gapSecs int)
}

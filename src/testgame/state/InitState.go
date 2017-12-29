package state

import (
	"../../base/ginterface"
	"../../base/nodestate"
)

// InitState is the initial state
type InitState struct {
	*nodestate.NodeState // base class
}

// OnEnter is called when entering this state
func (s *InitState) OnEnter() {

}

// OnExit is called when exiting this state
func (s *InitState) OnExit() {

}

// OnUpdate is called when updating this state
func (s *InitState) OnUpdate(gapSecs int) {

}

// NewInitState is a constructor of InitState
func NewInitState(node ginterface.INode) *InitState {
	ret := &InitState{}
	ret.NodeState = nodestate.NewNodeState(node)
	return ret
}

package state

import (
	"../../base/ginterface"
	"../../base/nodestate"
)

// DerivedState is the custom state of this state
type DerivedState struct {
	*nodestate.NodeState // base class
}

// OnEnter is called when entering this state
func (s *DerivedState) OnEnter() {

}

// OnExit is called when exiting this state
func (s *DerivedState) OnExit() {

}

// OnUpdate is called when updating this state
func (s *DerivedState) OnUpdate(gapSecs int) {

}

// NewDerivedState is a constructor of DerivedState
func NewDerivedState(node ginterface.INode) *DerivedState {
	ret := &DerivedState{}
	ret.NodeState = nodestate.NewNodeState(node)
	return ret
}

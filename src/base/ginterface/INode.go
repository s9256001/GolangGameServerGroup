package ginterface

// INode is an interface of a node companied with model and state
type INode interface {
	// GetModel gets the model
	GetModel() INodeModel
	// GetState gets current state
	GetState() INodeState
	// ChangeState changes the state of the node
	ChangeState(newState INodeState)
	// OnUpdate is called when updating this node
	OnUpdate(gapSecs int)
}

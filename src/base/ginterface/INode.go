package ginterface

// INode is an interface of a node companied with model and state
type INode interface {
	// GetLogger returns the game logger
	GetLogger() IGameLogger
	// GetModule returns the specific module to resolve import cycle
	GetModule(module interface{}) interface{}
	// GetModel gets the model
	GetModel() INodeModel
	// GetState gets current state
	GetState() INodeState
	// ChangeState changes the state of the node
	ChangeState(newState INodeState)
	// OnUpdate is called when updating this node
	OnUpdate(gapSecs int)
}

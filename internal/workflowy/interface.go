package workflowy

// WorkflowyClient defines the interface for Workflowy API operations
type WorkflowyClient interface {
	// Core CRUD operations
	CreateNode(req *CreateNodeRequest) (*CreateNodeResponse, error)
	GetNode(nodeID string) (*Node, error)
	UpdateNode(nodeID string, req *UpdateNodeRequest) (*StatusResponse, error)
	DeleteNode(nodeID string) (*StatusResponse, error)
	ListNodes(parentID string) ([]Node, error)

	// Node completion operations
	CompleteNode(nodeID string) (*StatusResponse, error)
	UncompleteNode(nodeID string) (*StatusResponse, error)

	// Helper methods
	GetTopLevelNodes() ([]Node, error)
	GetChildNodes(parentID string) ([]Node, error)
}

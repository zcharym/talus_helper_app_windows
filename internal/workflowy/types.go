package workflowy

// Node represents a WorkFlowy node (bullet point)
type Node struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Note        *string  `json:"note"`
	Priority    int      `json:"priority"`
	Data        NodeData `json:"data"`
	CreatedAt   int64    `json:"createdAt"`
	ModifiedAt  int64    `json:"modifiedAt"`
	CompletedAt *int64   `json:"completedAt"`
}

// NodeData contains additional node metadata
type NodeData struct {
	LayoutMode string `json:"layoutMode"`
}

// CreateNodeRequest represents the request payload for creating a node
type CreateNodeRequest struct {
	ParentID   string `json:"parent_id"`
	Name       string `json:"name"`
	Note       string `json:"note,omitempty"`
	LayoutMode string `json:"layoutMode,omitempty"`
	Position   string `json:"position,omitempty"`
}

// UpdateNodeRequest represents the request payload for updating a node
type UpdateNodeRequest struct {
	Name       string `json:"name,omitempty"`
	Note       string `json:"note,omitempty"`
	LayoutMode string `json:"layoutMode,omitempty"`
}

// CreateNodeResponse represents the response from creating a node
type CreateNodeResponse struct {
	ItemID string `json:"item_id"`
}

// GetNodeResponse represents the response from getting a node
type GetNodeResponse struct {
	Node Node `json:"node"`
}

// ListNodesResponse represents the response from listing nodes
type ListNodesResponse struct {
	Nodes []Node `json:"nodes"`
}

// StatusResponse represents a simple status response
type StatusResponse struct {
	Status string `json:"status"`
}

// ClientConfig holds configuration for the WorkFlowy API client
type ClientConfig struct {
	APIKey  string
	BaseURL string
}

// NewClientConfig creates a new client configuration with default values
func NewClientConfig(apiKey string) *ClientConfig {
	return &ClientConfig{
		APIKey:  apiKey,
		BaseURL: "https://workflowy.com/api/v1",
	}
}

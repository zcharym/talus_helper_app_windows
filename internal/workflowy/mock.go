package workflowy

import (
	"fmt"
	"sync"
	"time"
)

// MockClient is a mock implementation of WorkflowyClient for testing
type MockClient struct {
	// In-memory storage
	nodes map[string]*Node
	// Call tracking for testing
	CreateNodeCalls     []*CreateNodeRequest
	GetNodeCalls        []string
	UpdateNodeCalls     map[string]*UpdateNodeRequest
	DeleteNodeCalls     []string
	ListNodesCalls      []string
	CompleteNodeCalls   []string
	UncompleteNodeCalls []string

	// Error simulation
	ShouldErrorOnCreate     bool
	ShouldErrorOnGet        bool
	ShouldErrorOnUpdate     bool
	ShouldErrorOnDelete     bool
	ShouldErrorOnList       bool
	ShouldErrorOnComplete   bool
	ShouldErrorOnUncomplete bool

	// Custom error messages
	ErrorMessages map[string]string

	// Mutex for thread safety
	mu sync.RWMutex
}

// NewMockClient creates a new mock client
func NewMockClient() *MockClient {
	return &MockClient{
		nodes:           make(map[string]*Node),
		UpdateNodeCalls: make(map[string]*UpdateNodeRequest),
		ErrorMessages:   make(map[string]string),
	}
}

// CreateNode creates a new node in the mock
func (m *MockClient) CreateNode(req *CreateNodeRequest) (*CreateNodeResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.CreateNodeCalls = append(m.CreateNodeCalls, req)

	if m.ShouldErrorOnCreate {
		errorMsg := "mock create error"
		if msg, exists := m.ErrorMessages["create"]; exists {
			errorMsg = msg
		}
		return nil, fmt.Errorf("%s", errorMsg)
	}

	// Generate a mock ID with additional randomness
	nodeID := fmt.Sprintf("mock_%d_%d", time.Now().UnixNano(), len(m.nodes))

	// Create the node
	node := &Node{
		ID:          nodeID,
		Name:        req.Name,
		Note:        &req.Note,
		Priority:    0,
		Data:        NodeData{LayoutMode: req.LayoutMode},
		CreatedAt:   time.Now().Unix(),
		ModifiedAt:  time.Now().Unix(),
		CompletedAt: nil,
	}

	// Store the node
	m.nodes[nodeID] = node

	return &CreateNodeResponse{ItemID: nodeID}, nil
}

// GetNode retrieves a node by ID
func (m *MockClient) GetNode(nodeID string) (*Node, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.GetNodeCalls = append(m.GetNodeCalls, nodeID)

	if m.ShouldErrorOnGet {
		errorMsg := "mock get error"
		if msg, exists := m.ErrorMessages["get"]; exists {
			errorMsg = msg
		}
		return nil, fmt.Errorf("%s", errorMsg)
	}

	node, exists := m.nodes[nodeID]
	if !exists {
		return nil, fmt.Errorf("node not found")
	}

	// Return a copy to avoid race conditions
	nodeCopy := *node
	return &nodeCopy, nil
}

// UpdateNode updates an existing node
func (m *MockClient) UpdateNode(nodeID string, req *UpdateNodeRequest) (*StatusResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.UpdateNodeCalls[nodeID] = req

	if m.ShouldErrorOnUpdate {
		errorMsg := "mock update error"
		if msg, exists := m.ErrorMessages["update"]; exists {
			errorMsg = msg
		}
		return nil, fmt.Errorf("%s", errorMsg)
	}

	node, exists := m.nodes[nodeID]
	if !exists {
		return nil, fmt.Errorf("node not found")
	}

	// Update the node
	if req.Name != "" {
		node.Name = req.Name
	}
	if req.Note != "" {
		node.Note = &req.Note
	}
	if req.LayoutMode != "" {
		node.Data.LayoutMode = req.LayoutMode
	}
	node.ModifiedAt = time.Now().Unix()

	return &StatusResponse{Status: "success"}, nil
}

// DeleteNode deletes a node
func (m *MockClient) DeleteNode(nodeID string) (*StatusResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.DeleteNodeCalls = append(m.DeleteNodeCalls, nodeID)

	if m.ShouldErrorOnDelete {
		errorMsg := "mock delete error"
		if msg, exists := m.ErrorMessages["delete"]; exists {
			errorMsg = msg
		}
		return nil, fmt.Errorf("%s", errorMsg)
	}

	_, exists := m.nodes[nodeID]
	if !exists {
		return nil, fmt.Errorf("node not found")
	}

	delete(m.nodes, nodeID)
	return &StatusResponse{Status: "success"}, nil
}

// ListNodes retrieves child nodes for a given parent
func (m *MockClient) ListNodes(parentID string) ([]Node, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	m.ListNodesCalls = append(m.ListNodesCalls, parentID)

	if m.ShouldErrorOnList {
		errorMsg := "mock list error"
		if msg, exists := m.ErrorMessages["list"]; exists {
			errorMsg = msg
		}
		return nil, fmt.Errorf("%s", errorMsg)
	}

	var nodes []Node
	for _, node := range m.nodes {
		// In a real implementation, you'd filter by parent_id
		// For the mock, we'll return all nodes
		nodes = append(nodes, *node)
	}

	return nodes, nil
}

// CompleteNode marks a node as completed
func (m *MockClient) CompleteNode(nodeID string) (*StatusResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.CompleteNodeCalls = append(m.CompleteNodeCalls, nodeID)

	if m.ShouldErrorOnComplete {
		errorMsg := "mock complete error"
		if msg, exists := m.ErrorMessages["complete"]; exists {
			errorMsg = msg
		}
		return nil, fmt.Errorf("%s", errorMsg)
	}

	node, exists := m.nodes[nodeID]
	if !exists {
		return nil, fmt.Errorf("node not found")
	}

	now := time.Now().Unix()
	node.CompletedAt = &now
	node.ModifiedAt = now

	return &StatusResponse{Status: "success"}, nil
}

// UncompleteNode marks a node as not completed
func (m *MockClient) UncompleteNode(nodeID string) (*StatusResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.UncompleteNodeCalls = append(m.UncompleteNodeCalls, nodeID)

	if m.ShouldErrorOnUncomplete {
		errorMsg := "mock uncomplete error"
		if msg, exists := m.ErrorMessages["uncomplete"]; exists {
			errorMsg = msg
		}
		return nil, fmt.Errorf("%s", errorMsg)
	}

	node, exists := m.nodes[nodeID]
	if !exists {
		return nil, fmt.Errorf("node not found")
	}

	node.CompletedAt = nil
	node.ModifiedAt = time.Now().Unix()

	return &StatusResponse{Status: "success"}, nil
}

// GetTopLevelNodes retrieves all top-level nodes
func (m *MockClient) GetTopLevelNodes() ([]Node, error) {
	return m.ListNodes("None")
}

// GetChildNodes retrieves child nodes for a given parent
func (m *MockClient) GetChildNodes(parentID string) ([]Node, error) {
	return m.ListNodes(parentID)
}

// Helper methods for testing

// Reset clears all mock data and call tracking
func (m *MockClient) Reset() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.nodes = make(map[string]*Node)
	m.CreateNodeCalls = nil
	m.GetNodeCalls = nil
	m.UpdateNodeCalls = make(map[string]*UpdateNodeRequest)
	m.DeleteNodeCalls = nil
	m.ListNodesCalls = nil
	m.CompleteNodeCalls = nil
	m.UncompleteNodeCalls = nil

	m.ShouldErrorOnCreate = false
	m.ShouldErrorOnGet = false
	m.ShouldErrorOnUpdate = false
	m.ShouldErrorOnDelete = false
	m.ShouldErrorOnList = false
	m.ShouldErrorOnComplete = false
	m.ShouldErrorOnUncomplete = false

	m.ErrorMessages = make(map[string]string)
}

// AddNode adds a node to the mock for testing
func (m *MockClient) AddNode(node *Node) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.nodes[node.ID] = node
}

// GetNodeCount returns the number of nodes in the mock
func (m *MockClient) GetNodeCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.nodes)
}

// SetError simulates an error for a specific operation
func (m *MockClient) SetError(operation string, shouldError bool, message string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	switch operation {
	case "create":
		m.ShouldErrorOnCreate = shouldError
	case "get":
		m.ShouldErrorOnGet = shouldError
	case "update":
		m.ShouldErrorOnUpdate = shouldError
	case "delete":
		m.ShouldErrorOnDelete = shouldError
	case "list":
		m.ShouldErrorOnList = shouldError
	case "complete":
		m.ShouldErrorOnComplete = shouldError
	case "uncomplete":
		m.ShouldErrorOnUncomplete = shouldError
	}

	if message != "" {
		m.ErrorMessages[operation] = message
	}
}

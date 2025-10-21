package workflowy

import (
	"fmt"
	"testing"
	"time"
)

func TestMockClient_CreateNode(t *testing.T) {
	mock := NewMockClient()

	t.Run("successful creation", func(t *testing.T) {
		req := &CreateNodeRequest{
			ParentID:   "parent123",
			Name:       "Test Node",
			Note:       "Test note",
			LayoutMode: "list",
		}

		resp, err := mock.CreateNode(req)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if resp.ItemID == "" {
			t.Error("Expected ItemID to be set")
		}

		// Verify the node was stored
		if mock.GetNodeCount() != 1 {
			t.Errorf("Expected 1 node, got %d", mock.GetNodeCount())
		}

		// Verify call tracking
		if len(mock.CreateNodeCalls) != 1 {
			t.Errorf("Expected 1 create call, got %d", len(mock.CreateNodeCalls))
		}

		if mock.CreateNodeCalls[0].Name != req.Name {
			t.Errorf("Expected name %s, got %s", req.Name, mock.CreateNodeCalls[0].Name)
		}
	})

	t.Run("error simulation", func(t *testing.T) {
		mock.Reset()
		mock.SetError("create", true, "test error")

		req := &CreateNodeRequest{
			Name: "Test Node",
		}

		_, err := mock.CreateNode(req)
		if err == nil {
			t.Error("Expected error, got nil")
		}

		if err.Error() != "test error" {
			t.Errorf("Expected 'test error', got %s", err.Error())
		}
	})
}

func TestMockClient_GetNode(t *testing.T) {
	mock := NewMockClient()

	t.Run("successful retrieval", func(t *testing.T) {
		// Add a test node
		testNode := &Node{
			ID:        "test123",
			Name:      "Test Node",
			Priority:  1,
			CreatedAt: time.Now().Unix(),
		}
		mock.AddNode(testNode)

		node, err := mock.GetNode("test123")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if node.ID != "test123" {
			t.Errorf("Expected ID 'test123', got %s", node.ID)
		}

		if node.Name != "Test Node" {
			t.Errorf("Expected name 'Test Node', got %s", node.Name)
		}

		// Verify call tracking
		if len(mock.GetNodeCalls) != 1 {
			t.Errorf("Expected 1 get call, got %d", len(mock.GetNodeCalls))
		}
	})

	t.Run("node not found", func(t *testing.T) {
		mock.Reset()

		_, err := mock.GetNode("nonexistent")
		if err == nil {
			t.Error("Expected error for nonexistent node, got nil")
		}

		if err.Error() != "node not found" {
			t.Errorf("Expected 'node not found', got %s", err.Error())
		}
	})

	t.Run("error simulation", func(t *testing.T) {
		mock.Reset()
		mock.SetError("get", true, "test get error")

		_, err := mock.GetNode("test123")
		if err == nil {
			t.Error("Expected error, got nil")
		}

		if err.Error() != "test get error" {
			t.Errorf("Expected 'test get error', got %s", err.Error())
		}
	})
}

func TestMockClient_UpdateNode(t *testing.T) {
	mock := NewMockClient()

	t.Run("successful update", func(t *testing.T) {
		// Add a test node
		testNode := &Node{
			ID:        "test123",
			Name:      "Original Name",
			Priority:  1,
			CreatedAt: time.Now().Unix(),
		}
		mock.AddNode(testNode)

		req := &UpdateNodeRequest{
			Name:       "Updated Name",
			Note:       "Updated note",
			LayoutMode: "board",
		}

		resp, err := mock.UpdateNode("test123", req)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if resp.Status != "success" {
			t.Errorf("Expected status 'success', got %s", resp.Status)
		}

		// Verify the node was updated
		node, err := mock.GetNode("test123")
		if err != nil {
			t.Fatalf("Failed to get updated node: %v", err)
		}

		if node.Name != "Updated Name" {
			t.Errorf("Expected name 'Updated Name', got %s", node.Name)
		}

		if node.Data.LayoutMode != "board" {
			t.Errorf("Expected layout mode 'board', got %s", node.Data.LayoutMode)
		}

		// Verify call tracking
		if len(mock.UpdateNodeCalls) != 1 {
			t.Errorf("Expected 1 update call, got %d", len(mock.UpdateNodeCalls))
		}
	})

	t.Run("node not found", func(t *testing.T) {
		mock.Reset()

		req := &UpdateNodeRequest{Name: "Test"}
		_, err := mock.UpdateNode("nonexistent", req)
		if err == nil {
			t.Error("Expected error for nonexistent node, got nil")
		}
	})

	t.Run("error simulation", func(t *testing.T) {
		mock.Reset()
		mock.SetError("update", true, "test update error")

		req := &UpdateNodeRequest{Name: "Test"}
		_, err := mock.UpdateNode("test123", req)
		if err == nil {
			t.Error("Expected error, got nil")
		}

		if err.Error() != "test update error" {
			t.Errorf("Expected 'test update error', got %s", err.Error())
		}
	})
}

func TestMockClient_DeleteNode(t *testing.T) {
	mock := NewMockClient()

	t.Run("successful deletion", func(t *testing.T) {
		// Add a test node
		testNode := &Node{
			ID:   "test123",
			Name: "Test Node",
		}
		mock.AddNode(testNode)

		if mock.GetNodeCount() != 1 {
			t.Errorf("Expected 1 node before deletion, got %d", mock.GetNodeCount())
		}

		resp, err := mock.DeleteNode("test123")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if resp.Status != "success" {
			t.Errorf("Expected status 'success', got %s", resp.Status)
		}

		// Verify the node was deleted
		if mock.GetNodeCount() != 0 {
			t.Errorf("Expected 0 nodes after deletion, got %d", mock.GetNodeCount())
		}

		// Verify call tracking
		if len(mock.DeleteNodeCalls) != 1 {
			t.Errorf("Expected 1 delete call, got %d", len(mock.DeleteNodeCalls))
		}
	})

	t.Run("node not found", func(t *testing.T) {
		mock.Reset()

		_, err := mock.DeleteNode("nonexistent")
		if err == nil {
			t.Error("Expected error for nonexistent node, got nil")
		}
	})

	t.Run("error simulation", func(t *testing.T) {
		mock.Reset()
		mock.SetError("delete", true, "test delete error")

		_, err := mock.DeleteNode("test123")
		if err == nil {
			t.Error("Expected error, got nil")
		}

		if err.Error() != "test delete error" {
			t.Errorf("Expected 'test delete error', got %s", err.Error())
		}
	})
}

func TestMockClient_ListNodes(t *testing.T) {
	mock := NewMockClient()

	t.Run("successful listing", func(t *testing.T) {
		// Add test nodes
		node1 := &Node{ID: "node1", Name: "Node 1"}
		node2 := &Node{ID: "node2", Name: "Node 2"}
		mock.AddNode(node1)
		mock.AddNode(node2)

		nodes, err := mock.ListNodes("parent123")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if len(nodes) != 2 {
			t.Errorf("Expected 2 nodes, got %d", len(nodes))
		}

		// Verify call tracking
		if len(mock.ListNodesCalls) != 1 {
			t.Errorf("Expected 1 list call, got %d", len(mock.ListNodesCalls))
		}

		if mock.ListNodesCalls[0] != "parent123" {
			t.Errorf("Expected parent_id 'parent123', got %s", mock.ListNodesCalls[0])
		}
	})

	t.Run("empty list", func(t *testing.T) {
		mock.Reset()

		nodes, err := mock.ListNodes("parent123")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if len(nodes) != 0 {
			t.Errorf("Expected 0 nodes, got %d", len(nodes))
		}
	})

	t.Run("error simulation", func(t *testing.T) {
		mock.Reset()
		mock.SetError("list", true, "test list error")

		_, err := mock.ListNodes("parent123")
		if err == nil {
			t.Error("Expected error, got nil")
		}

		if err.Error() != "test list error" {
			t.Errorf("Expected 'test list error', got %s", err.Error())
		}
	})
}

func TestMockClient_CompleteNode(t *testing.T) {
	mock := NewMockClient()

	t.Run("successful completion", func(t *testing.T) {
		// Add a test node
		testNode := &Node{
			ID:        "test123",
			Name:      "Test Node",
			CreatedAt: time.Now().Unix(),
		}
		mock.AddNode(testNode)

		resp, err := mock.CompleteNode("test123")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if resp.Status != "success" {
			t.Errorf("Expected status 'success', got %s", resp.Status)
		}

		// Verify the node was marked as completed
		node, err := mock.GetNode("test123")
		if err != nil {
			t.Fatalf("Failed to get completed node: %v", err)
		}

		if node.CompletedAt == nil {
			t.Error("Expected node to be completed")
		}

		// Verify call tracking
		if len(mock.CompleteNodeCalls) != 1 {
			t.Errorf("Expected 1 complete call, got %d", len(mock.CompleteNodeCalls))
		}
	})

	t.Run("node not found", func(t *testing.T) {
		mock.Reset()

		_, err := mock.CompleteNode("nonexistent")
		if err == nil {
			t.Error("Expected error for nonexistent node, got nil")
		}
	})

	t.Run("error simulation", func(t *testing.T) {
		mock.Reset()
		mock.SetError("complete", true, "test complete error")

		_, err := mock.CompleteNode("test123")
		if err == nil {
			t.Error("Expected error, got nil")
		}

		if err.Error() != "test complete error" {
			t.Errorf("Expected 'test complete error', got %s", err.Error())
		}
	})
}

func TestMockClient_UncompleteNode(t *testing.T) {
	mock := NewMockClient()

	t.Run("successful uncompletion", func(t *testing.T) {
		// Add a completed test node
		now := time.Now().Unix()
		testNode := &Node{
			ID:          "test123",
			Name:        "Test Node",
			CompletedAt: &now,
			CreatedAt:   time.Now().Unix(),
		}
		mock.AddNode(testNode)

		resp, err := mock.UncompleteNode("test123")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if resp.Status != "success" {
			t.Errorf("Expected status 'success', got %s", resp.Status)
		}

		// Verify the node was marked as not completed
		node, err := mock.GetNode("test123")
		if err != nil {
			t.Fatalf("Failed to get uncompleted node: %v", err)
		}

		if node.CompletedAt != nil {
			t.Error("Expected node to be uncompleted")
		}

		// Verify call tracking
		if len(mock.UncompleteNodeCalls) != 1 {
			t.Errorf("Expected 1 uncomplete call, got %d", len(mock.UncompleteNodeCalls))
		}
	})

	t.Run("node not found", func(t *testing.T) {
		mock.Reset()

		_, err := mock.UncompleteNode("nonexistent")
		if err == nil {
			t.Error("Expected error for nonexistent node, got nil")
		}
	})

	t.Run("error simulation", func(t *testing.T) {
		mock.Reset()
		mock.SetError("uncomplete", true, "test uncomplete error")

		_, err := mock.UncompleteNode("test123")
		if err == nil {
			t.Error("Expected error, got nil")
		}

		if err.Error() != "test uncomplete error" {
			t.Errorf("Expected 'test uncomplete error', got %s", err.Error())
		}
	})
}

func TestMockClient_HelperMethods(t *testing.T) {
	mock := NewMockClient()

	t.Run("GetTopLevelNodes", func(t *testing.T) {
		// Add test nodes
		node1 := &Node{ID: "node1", Name: "Node 1"}
		node2 := &Node{ID: "node2", Name: "Node 2"}
		mock.AddNode(node1)
		mock.AddNode(node2)

		nodes, err := mock.GetTopLevelNodes()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if len(nodes) != 2 {
			t.Errorf("Expected 2 nodes, got %d", len(nodes))
		}

		// Verify it called ListNodes with "None"
		if len(mock.ListNodesCalls) != 1 {
			t.Errorf("Expected 1 list call, got %d", len(mock.ListNodesCalls))
		}

		if mock.ListNodesCalls[0] != "None" {
			t.Errorf("Expected parent_id 'None', got %s", mock.ListNodesCalls[0])
		}
	})

	t.Run("GetChildNodes", func(t *testing.T) {
		mock.Reset()

		// Add test nodes
		node1 := &Node{ID: "node1", Name: "Node 1"}
		mock.AddNode(node1)

		nodes, err := mock.GetChildNodes("parent123")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if len(nodes) != 1 {
			t.Errorf("Expected 1 node, got %d", len(nodes))
		}

		// Verify it called ListNodes with the parent ID
		if len(mock.ListNodesCalls) != 1 {
			t.Errorf("Expected 1 list call, got %d", len(mock.ListNodesCalls))
		}

		if mock.ListNodesCalls[0] != "parent123" {
			t.Errorf("Expected parent_id 'parent123', got %s", mock.ListNodesCalls[0])
		}
	})
}

func TestMockClient_Reset(t *testing.T) {
	mock := NewMockClient()

	// Add some data and make some calls
	testNode := &Node{ID: "test123", Name: "Test Node"}
	mock.AddNode(testNode)

	_, _ = mock.CreateNode(&CreateNodeRequest{Name: "Test"})
	_, _ = mock.GetNode("test123")
	_, _ = mock.ListNodes("parent")

	// Set some errors
	mock.SetError("create", true, "test error")

	// Reset
	mock.Reset()

	// Verify everything is cleared
	if mock.GetNodeCount() != 0 {
		t.Errorf("Expected 0 nodes after reset, got %d", mock.GetNodeCount())
	}

	if len(mock.CreateNodeCalls) != 0 {
		t.Errorf("Expected 0 create calls after reset, got %d", len(mock.CreateNodeCalls))
	}

	if len(mock.GetNodeCalls) != 0 {
		t.Errorf("Expected 0 get calls after reset, got %d", len(mock.GetNodeCalls))
	}

	if len(mock.ListNodesCalls) != 0 {
		t.Errorf("Expected 0 list calls after reset, got %d", len(mock.ListNodesCalls))
	}

	if mock.ShouldErrorOnCreate {
		t.Error("Expected error flags to be cleared after reset")
	}
}

func TestMockClient_Concurrency(t *testing.T) {
	mock := NewMockClient()

	// Test concurrent access
	done := make(chan bool, 10)
	errors := make(chan error, 10)

	for i := 0; i < 10; i++ {
		go func(id int) {
			defer func() { done <- true }()

			// Create a node with unique name
			req := &CreateNodeRequest{
				Name:     fmt.Sprintf("Concurrent Node %d", id),
				ParentID: "parent",
			}
			_, err := mock.CreateNode(req)
			if err != nil {
				errors <- fmt.Errorf("goroutine %d error creating node: %v", id, err)
			}
		}(i)
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}

	// Check for errors
	close(errors)
	for err := range errors {
		t.Error(err)
	}

	// Verify we have 10 nodes
	nodeCount := mock.GetNodeCount()
	if nodeCount != 10 {
		t.Errorf("Expected 10 nodes, got %d", nodeCount)
	}

	// Verify we have 10 create calls
	callCount := len(mock.CreateNodeCalls)
	if callCount != 10 {
		t.Errorf("Expected 10 create calls, got %d", callCount)
	}
}

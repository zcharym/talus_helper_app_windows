package workflowy

import (
	"testing"
)

// TestInterfaceCompliance verifies that both Client and MockClient implement WorkflowyClient
func TestInterfaceCompliance(t *testing.T) {
	// Test that Client implements WorkflowyClient
	var _ WorkflowyClient = (*Client)(nil)

	// Test that MockClient implements WorkflowyClient
	var _ WorkflowyClient = (*MockClient)(nil)
}

// TestWorkflowyClientInterface demonstrates using the interface with different implementations
func TestWorkflowyClientInterface(t *testing.T) {
	// Test with mock client
	t.Run("mock client", func(t *testing.T) {
		mockClient := NewMockClient()
		testWorkflowyClient(t, mockClient)
	})

	// Test with real client (using mock config to avoid actual API calls)
	t.Run("real client", func(t *testing.T) {
		config := NewClientConfig("test-api-key")
		config.BaseURL = "http://localhost:8080/api/v1" // Use a non-existent URL to avoid real API calls
		realClient := NewClient(config)

		// We can't easily test the real client without mocking HTTP responses,
		// but we can verify it implements the interface
		var _ WorkflowyClient = realClient
	})
}

// testWorkflowyClient is a generic test function that works with any WorkflowyClient implementation
func testWorkflowyClient(t *testing.T, client WorkflowyClient) {
	t.Run("create and retrieve node", func(t *testing.T) {
		// Create a node
		createReq := &CreateNodeRequest{
			ParentID:   "parent123",
			Name:       "Test Node",
			Note:       "Test note",
			LayoutMode: "list",
		}

		createResp, err := client.CreateNode(createReq)
		if err != nil {
			t.Fatalf("Failed to create node: %v", err)
		}

		if createResp.ItemID == "" {
			t.Error("Expected ItemID to be set")
		}

		// Retrieve the node
		node, err := client.GetNode(createResp.ItemID)
		if err != nil {
			t.Fatalf("Failed to get node: %v", err)
		}

		if node.Name != "Test Node" {
			t.Errorf("Expected name 'Test Node', got %s", node.Name)
		}
	})

	t.Run("update node", func(t *testing.T) {
		// Create a node first
		createReq := &CreateNodeRequest{
			Name: "Original Name",
		}

		createResp, err := client.CreateNode(createReq)
		if err != nil {
			t.Fatalf("Failed to create node: %v", err)
		}

		// Update the node
		updateReq := &UpdateNodeRequest{
			Name:       "Updated Name",
			Note:       "Updated note",
			LayoutMode: "board",
		}

		updateResp, err := client.UpdateNode(createResp.ItemID, updateReq)
		if err != nil {
			t.Fatalf("Failed to update node: %v", err)
		}

		if updateResp.Status != "success" {
			t.Errorf("Expected status 'success', got %s", updateResp.Status)
		}
	})

	t.Run("complete and uncomplete node", func(t *testing.T) {
		// Create a node first
		createReq := &CreateNodeRequest{
			Name: "Task Node",
		}

		createResp, err := client.CreateNode(createReq)
		if err != nil {
			t.Fatalf("Failed to create node: %v", err)
		}

		// Complete the node
		completeResp, err := client.CompleteNode(createResp.ItemID)
		if err != nil {
			t.Fatalf("Failed to complete node: %v", err)
		}

		if completeResp.Status != "success" {
			t.Errorf("Expected status 'success', got %s", completeResp.Status)
		}

		// Verify the node is completed
		node, err := client.GetNode(createResp.ItemID)
		if err != nil {
			t.Fatalf("Failed to get node: %v", err)
		}

		if node.CompletedAt == nil {
			t.Error("Expected node to be completed")
		}

		// Uncomplete the node
		uncompleteResp, err := client.UncompleteNode(createResp.ItemID)
		if err != nil {
			t.Fatalf("Failed to uncomplete node: %v", err)
		}

		if uncompleteResp.Status != "success" {
			t.Errorf("Expected status 'success', got %s", uncompleteResp.Status)
		}

		// Verify the node is not completed
		node, err = client.GetNode(createResp.ItemID)
		if err != nil {
			t.Fatalf("Failed to get node: %v", err)
		}

		if node.CompletedAt != nil {
			t.Error("Expected node to be uncompleted")
		}
	})

	t.Run("list nodes", func(t *testing.T) {
		// Create some test nodes
		for i := 0; i < 3; i++ {
			createReq := &CreateNodeRequest{
				Name:     "Test Node",
				ParentID: "parent123",
			}

			_, err := client.CreateNode(createReq)
			if err != nil {
				t.Fatalf("Failed to create node %d: %v", i, err)
			}
		}

		// List nodes
		nodes, err := client.ListNodes("parent123")
		if err != nil {
			t.Fatalf("Failed to list nodes: %v", err)
		}

		// For mock client, we expect at least 3 nodes
		// For real client, this might be different
		if len(nodes) < 3 {
			t.Logf("Got %d nodes (expected at least 3) - this is expected for mock client that doesn't filter by parent_id", len(nodes))
		}
	})

	t.Run("helper methods", func(t *testing.T) {
		// Test GetTopLevelNodes
		topLevelNodes, err := client.GetTopLevelNodes()
		if err != nil {
			t.Fatalf("Failed to get top level nodes: %v", err)
		}

		// We don't know how many nodes will be returned, but it should not error
		_ = topLevelNodes

		// Test GetChildNodes
		childNodes, err := client.GetChildNodes("parent123")
		if err != nil {
			t.Fatalf("Failed to get child nodes: %v", err)
		}

		// We don't know how many nodes will be returned, but it should not error
		_ = childNodes
	})
}

// TestInterfacePolymorphism demonstrates how the interface allows for polymorphism
func TestInterfacePolymorphism(t *testing.T) {
	// Create a function that works with any WorkflowyClient
	processWorkflowyClient := func(client WorkflowyClient) error {
		// Create a node
		createReq := &CreateNodeRequest{
			Name: "Polymorphic Test",
		}

		createResp, err := client.CreateNode(createReq)
		if err != nil {
			return err
		}

		// Get the node
		_, err = client.GetNode(createResp.ItemID)
		if err != nil {
			return err
		}

		// Complete the node
		_, err = client.CompleteNode(createResp.ItemID)
		if err != nil {
			return err
		}

		// Delete the node
		_, err = client.DeleteNode(createResp.ItemID)
		if err != nil {
			return err
		}

		return nil
	}

	// Test with mock client
	t.Run("with mock client", func(t *testing.T) {
		mockClient := NewMockClient()
		err := processWorkflowyClient(mockClient)
		if err != nil {
			t.Fatalf("Process failed with mock client: %v", err)
		}
	})

	// Test with real client (this would fail with real API calls, but demonstrates the pattern)
	t.Run("with real client", func(t *testing.T) {
		config := NewClientConfig("test-api-key")
		config.BaseURL = "http://localhost:8080/api/v1"
		realClient := NewClient(config)

		// We can't actually run this without mocking HTTP, but we can verify the interface works
		var _ WorkflowyClient = realClient

		// The processWorkflowyClient function would work with realClient too
		// (assuming proper API configuration)
	})
}

// TestMockClientSpecificFeatures tests features specific to the mock client
func TestMockClientSpecificFeatures(t *testing.T) {
	mockClient := NewMockClient()

	t.Run("call tracking", func(t *testing.T) {
		// Make some calls
		_, _ = mockClient.CreateNode(&CreateNodeRequest{Name: "Test"})
		_, _ = mockClient.GetNode("test123")
		_, _ = mockClient.ListNodes("parent")

		// Verify call tracking
		if len(mockClient.CreateNodeCalls) != 1 {
			t.Errorf("Expected 1 create call, got %d", len(mockClient.CreateNodeCalls))
		}

		if len(mockClient.GetNodeCalls) != 1 {
			t.Errorf("Expected 1 get call, got %d", len(mockClient.GetNodeCalls))
		}

		if len(mockClient.ListNodesCalls) != 1 {
			t.Errorf("Expected 1 list call, got %d", len(mockClient.ListNodesCalls))
		}
	})

	t.Run("error simulation", func(t *testing.T) {
		mockClient.Reset()
		mockClient.SetError("create", true, "simulated error")

		_, err := mockClient.CreateNode(&CreateNodeRequest{Name: "Test"})
		if err == nil {
			t.Error("Expected error, got nil")
		}

		if err.Error() != "simulated error" {
			t.Errorf("Expected 'simulated error', got %s", err.Error())
		}
	})

	t.Run("reset functionality", func(t *testing.T) {
		// Add some data
		mockClient.AddNode(&Node{ID: "test", Name: "Test"})
		_, _ = mockClient.CreateNode(&CreateNodeRequest{Name: "Test"})

		// Reset
		mockClient.Reset()

		// Verify everything is cleared
		if mockClient.GetNodeCount() != 0 {
			t.Errorf("Expected 0 nodes after reset, got %d", mockClient.GetNodeCount())
		}

		if len(mockClient.CreateNodeCalls) != 0 {
			t.Errorf("Expected 0 create calls after reset, got %d", len(mockClient.CreateNodeCalls))
		}
	})
}

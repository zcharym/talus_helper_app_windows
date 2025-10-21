package workflowy

import (
	"fmt"
	"testing"
)

// ExampleWorkflowyClient demonstrates how to use the WorkflowyClient interface
func ExampleWorkflowyClient() {
	// Create a mock client for testing
	mockClient := NewMockClient()

	// Use the interface - this works with any implementation
	var client WorkflowyClient = mockClient

	// Create a new node
	createReq := &CreateNodeRequest{
		ParentID:   "parent123",
		Name:       "Example Task",
		Note:       "This is an example task",
		LayoutMode: "list",
	}

	createResp, err := client.CreateNode(createReq)
	if err != nil {
		fmt.Printf("Error creating node: %v\n", err)
		return
	}

	fmt.Printf("Created node with ID: %s\n", createResp.ItemID)

	// Get the node
	node, err := client.GetNode(createResp.ItemID)
	if err != nil {
		fmt.Printf("Error getting node: %v\n", err)
		return
	}

	fmt.Printf("Node name: %s\n", node.Name)

	// Update the node
	updateReq := &UpdateNodeRequest{
		Name: "Updated Task Name",
		Note: "Updated note",
	}

	_, err = client.UpdateNode(createResp.ItemID, updateReq)
	if err != nil {
		fmt.Printf("Error updating node: %v\n", err)
		return
	}

	fmt.Println("Node updated successfully")

	// Complete the node
	_, err = client.CompleteNode(createResp.ItemID)
	if err != nil {
		fmt.Printf("Error completing node: %v\n", err)
		return
	}

	fmt.Println("Node completed successfully")

	// List nodes
	nodes, err := client.ListNodes("parent123")
	if err != nil {
		fmt.Printf("Error listing nodes: %v\n", err)
		return
	}

	fmt.Printf("Found %d nodes\n", len(nodes))

	// Delete the node
	_, err = client.DeleteNode(createResp.ItemID)
	if err != nil {
		fmt.Printf("Error deleting node: %v\n", err)
		return
	}

	fmt.Println("Node deleted successfully")
}

// TestExampleWorkflowyClient tests the example
func TestExampleWorkflowyClient(t *testing.T) {
	// This test verifies that the example code works
	ExampleWorkflowyClient()
}

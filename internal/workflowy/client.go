// check for /docs/workflowy-api-reference.md for more information
package workflowy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// Client represents a WorkFlowy API client
// It implements the WorkflowyClient interface
type Client struct {
	config     *ClientConfig
	httpClient *http.Client
}

// NewClient creates a new WorkFlowy API client
func NewClient(config *ClientConfig) *Client {
	return &Client{
		config: config,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// CreateNode creates a new node in WorkFlowy
func (c *Client) CreateNode(req *CreateNodeRequest) (*CreateNodeResponse, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.makeRequest("POST", "/nodes/", jsonData)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var createResp CreateNodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&createResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &createResp, nil
}

// UpdateNode updates an existing node
func (c *Client) UpdateNode(nodeID string, req *UpdateNodeRequest) (*StatusResponse, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.makeRequest("POST", fmt.Sprintf("/nodes/%s", nodeID), jsonData)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var statusResp StatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&statusResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &statusResp, nil
}

// GetNode retrieves a specific node by ID
func (c *Client) GetNode(nodeID string) (*Node, error) {
	resp, err := c.makeRequest("GET", fmt.Sprintf("/nodes/%s", nodeID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var getResp GetNodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&getResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &getResp.Node, nil
}

// ListNodes retrieves child nodes for a given parent
func (c *Client) ListNodes(parentID string) ([]Node, error) {
	params := url.Values{}
	if parentID != "" {
		params.Set("parent_id", parentID)
	}

	url := fmt.Sprintf("%s/nodes", c.config.BaseURL)
	if len(params) > 0 {
		url += "?" + params.Encode()
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.config.APIKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var listResp ListNodesResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return listResp.Nodes, nil
}

// DeleteNode permanently deletes a node
func (c *Client) DeleteNode(nodeID string) (*StatusResponse, error) {
	resp, err := c.makeRequest("DELETE", fmt.Sprintf("/nodes/%s", nodeID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var statusResp StatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&statusResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &statusResp, nil
}

// CompleteNode marks a node as completed
func (c *Client) CompleteNode(nodeID string) (*StatusResponse, error) {
	resp, err := c.makeRequest("POST", fmt.Sprintf("/nodes/%s/complete", nodeID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var statusResp StatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&statusResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &statusResp, nil
}

// UncompleteNode marks a node as not completed
func (c *Client) UncompleteNode(nodeID string) (*StatusResponse, error) {
	resp, err := c.makeRequest("POST", fmt.Sprintf("/nodes/%s/uncomplete", nodeID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var statusResp StatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&statusResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &statusResp, nil
}

// makeRequest is a helper method to make HTTP requests to the WorkFlowy API
func (c *Client) makeRequest(method, path string, body []byte) (*http.Response, error) {
	url := c.config.BaseURL + path

	var reqBody io.Reader
	if body != nil {
		reqBody = bytes.NewBuffer(body)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.config.APIKey)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	return resp, nil
}

// Helper methods for common operations

// GetTopLevelNodes retrieves all top-level nodes
func (c *Client) GetTopLevelNodes() ([]Node, error) {
	return c.ListNodes("None")
}

// GetChildNodes retrieves child nodes for a given parent
func (c *Client) GetChildNodes(parentID string) ([]Node, error) {
	return c.ListNodes(parentID)
}

// FormatNodeAsString returns a formatted string representation of a node
func FormatNodeAsString(node *Node) string {
	status := "incomplete"
	if node.CompletedAt != nil {
		status = "completed"
	}

	note := ""
	if node.Note != nil && *node.Note != "" {
		note = fmt.Sprintf(" (Note: %s)", *node.Note)
	}

	return fmt.Sprintf("ID: %s | %s | Priority: %d | Status: %s | Layout: %s%s",
		node.ID, node.Name, node.Priority, status, node.Data.LayoutMode, note)
}

// FormatNodesAsList returns a formatted list of nodes
func FormatNodesAsList(nodes []Node) string {
	if len(nodes) == 0 {
		return "No nodes found."
	}

	var result string
	for i, node := range nodes {
		result += fmt.Sprintf("%d. %s\n", i+1, FormatNodeAsString(&node))
	}
	return result
}

# Workflowy API Reference

## Overview

The Workflowy API allows developers to connect Workflowy to other applications and automate workflows. It supports basic CRUD (Create, Read, Update, Delete) operations for managing nodes (bullet points) in your Workflowy account.

## Base URL

```
https://workflowy.com/api/v1
```

## Authentication

To authenticate API requests, you'll need to generate an API key:

1. Visit the [Workflowy API Key page](https://workflowy.com/api-key)
2. Generate and copy your API key
3. Include this API key in your request headers:

```
Authorization: Bearer YOUR_API_KEY
Content-Type: application/json
```

## Data Models

### Node

A node represents a bullet point in Workflowy.

```json
{
  "id": "string",
  "name": "string",
  "note": "string (optional)",
  "priority": "integer",
  "data": {
    "layoutMode": "string"
  },
  "createdAt": "integer (timestamp)",
  "modifiedAt": "integer (timestamp)",
  "completedAt": "integer (timestamp, optional)"
}
```

### NodeData

Contains additional metadata for a node.

```json
{
  "layoutMode": "string"
}
```

## Endpoints

### Create Node

**Endpoint:** `POST /nodes/`

**Description:** Creates a new node in your Workflowy account.

**Request Body:**

```json
{
  "parent_id": "string",
  "name": "string",
  "note": "string (optional)",
  "layoutMode": "string (optional)",
  "position": "string (optional)"
}
```

**Response:**

```json
{
  "item_id": "string"
}
```

**Example Request:**

```http
POST /api/v1/nodes/ HTTP/1.1
Host: workflowy.com
Authorization: Bearer YOUR_API_KEY
Content-Type: application/json

{
  "parent_id": "1234567890abcdef",
  "name": "New task",
  "note": "This is a note for the task",
  "layoutMode": "list"
}
```

**Example Response:**

```json
{
  "item_id": "abcdef1234567890"
}
```

### Get Node

**Endpoint:** `GET /nodes/{id}`

**Description:** Retrieves a specific node by its ID.

**Parameters:**
- `id` (string): The ID of the node to retrieve

**Response:**

```json
{
  "node": {
    "id": "string",
    "name": "string",
    "note": "string (optional)",
    "priority": "integer",
    "data": {
      "layoutMode": "string"
    },
    "createdAt": "integer",
    "modifiedAt": "integer",
    "completedAt": "integer (optional)"
  }
}
```

**Example Request:**

```http
GET /api/v1/nodes/abcdef1234567890 HTTP/1.1
Host: workflowy.com
Authorization: Bearer YOUR_API_KEY
```

**Example Response:**

```json
{
  "node": {
    "id": "abcdef1234567890",
    "name": "New task",
    "note": "This is a note for the task",
    "priority": 0,
    "data": {
      "layoutMode": "list"
    },
    "createdAt": 1634567890,
    "modifiedAt": 1634567890,
    "completedAt": null
  }
}
```

### Update Node

**Endpoint:** `POST /nodes/{id}`

**Description:** Updates an existing node.

**Parameters:**
- `id` (string): The ID of the node to update

**Request Body:**

```json
{
  "name": "string (optional)",
  "note": "string (optional)",
  "layoutMode": "string (optional)"
}
```

**Response:**

```json
{
  "status": "string"
}
```

**Example Request:**

```http
POST /api/v1/nodes/abcdef1234567890 HTTP/1.1
Host: workflowy.com
Authorization: Bearer YOUR_API_KEY
Content-Type: application/json

{
  "name": "Updated task name",
  "note": "Updated note"
}
```

**Example Response:**

```json
{
  "status": "success"
}
```

### List Nodes

**Endpoint:** `GET /nodes`

**Description:** Retrieves child nodes for a given parent.

**Query Parameters:**
- `parent_id` (string, optional): The ID of the parent node. Use "None" for top-level nodes.

**Response:**

```json
{
  "nodes": [
    {
      "id": "string",
      "name": "string",
      "note": "string (optional)",
      "priority": "integer",
      "data": {
        "layoutMode": "string"
      },
      "createdAt": "integer",
      "modifiedAt": "integer",
      "completedAt": "integer (optional)"
    }
  ]
}
```

**Example Request:**

```http
GET /api/v1/nodes?parent_id=1234567890abcdef HTTP/1.1
Host: workflowy.com
Authorization: Bearer YOUR_API_KEY
```

**Example Response:**

```json
{
  "nodes": [
    {
      "id": "abcdef1234567890",
      "name": "Task 1",
      "note": "First task",
      "priority": 0,
      "data": {
        "layoutMode": "list"
      },
      "createdAt": 1634567890,
      "modifiedAt": 1634567890,
      "completedAt": null
    },
    {
      "id": "fedcba0987654321",
      "name": "Task 2",
      "note": "Second task",
      "priority": 1,
      "data": {
        "layoutMode": "list"
      },
      "createdAt": 1634567891,
      "modifiedAt": 1634567891,
      "completedAt": 1634567892
    }
  ]
}
```

### Delete Node

**Endpoint:** `DELETE /nodes/{id}`

**Description:** Permanently deletes a node.

**Parameters:**
- `id` (string): The ID of the node to delete

**Response:**

```json
{
  "status": "string"
}
```

**Example Request:**

```http
DELETE /api/v1/nodes/abcdef1234567890 HTTP/1.1
Host: workflowy.com
Authorization: Bearer YOUR_API_KEY
```

**Example Response:**

```json
{
  "status": "success"
}
```

### Complete Node

**Endpoint:** `POST /nodes/{id}/complete`

**Description:** Marks a node as completed.

**Parameters:**
- `id` (string): The ID of the node to complete

**Response:**

```json
{
  "status": "string"
}
```

**Example Request:**

```http
POST /api/v1/nodes/abcdef1234567890/complete HTTP/1.1
Host: workflowy.com
Authorization: Bearer YOUR_API_KEY
```

**Example Response:**

```json
{
  "status": "success"
}
```

### Uncomplete Node

**Endpoint:** `POST /nodes/{id}/uncomplete`

**Description:** Marks a node as not completed.

**Parameters:**
- `id` (string): The ID of the node to uncomplete

**Response:**

```json
{
  "status": "string"
}
```

**Example Request:**

```http
POST /api/v1/nodes/abcdef1234567890/uncomplete HTTP/1.1
Host: workflowy.com
Authorization: Bearer YOUR_API_KEY
```

**Example Response:**

```json
{
  "status": "success"
}
```

## Error Handling

The API returns standard HTTP status codes to indicate the success or failure of a request:

- **200 OK:** The request was successful
- **400 Bad Request:** The request was invalid or cannot be served
- **401 Unauthorized:** Authentication failed or user does not have permissions
- **404 Not Found:** The requested resource could not be found
- **500 Internal Server Error:** An error occurred on the server

### Error Response Format

```json
{
  "error": "string",
  "message": "string"
}
```

## Rate Limiting

The Workflowy API has rate limiting in place. Please refer to the official documentation for current rate limits.

## Common Use Cases

### Getting Top-Level Nodes

To retrieve all top-level nodes, use the list nodes endpoint with `parent_id` set to "None":

```http
GET /api/v1/nodes?parent_id=None HTTP/1.1
Host: workflowy.com
Authorization: Bearer YOUR_API_KEY
```

### Creating a Hierarchical Structure

1. Create a parent node
2. Use the returned `item_id` as the `parent_id` for child nodes
3. Repeat for deeper nesting levels

### Managing Task Completion

- Use the complete endpoint to mark tasks as done
- Use the uncomplete endpoint to mark tasks as not done
- Check the `completedAt` field in node data to determine completion status

## SDK and Libraries

This API reference is based on the Go implementation found in the codebase. For other languages, you can use the same endpoints with appropriate HTTP libraries.

## Additional Resources

- [Workflowy Integrations Documentation](https://workflowy.com/learn/integrations/)
- [Workflowy API Key Generation](https://workflowy.com/api-key)
- [Workflowy Support](https://workflowy.com/learn/)

## Changelog

- **v1.0.0:** Initial API reference based on current implementation
- Last updated: 2024-01-21

package golucky

import "fmt"

type Group struct {
	Name        string   `json:"name"`
	DisplayName string   `json:"displayName"`
	Weight      int      `json:"weight"`
	Nodes       []Node   `json:"nodes"`
	Metadata    Metadata `json:"metadata"`
}

type groupName struct {
	Name string `json:"name"`
}

func (c *Client) ListGroups() (*[]string, error) {
	return getRequest[[]string](fmt.Sprintf("%s/group", c.config.RestIp), c.config.AuthKey)
}

func (c *Client) CreateGroup(name string) (*Group, error) {
	return postRequestBody[Group](fmt.Sprintf("%s/group", c.config.RestIp), groupName{Name: name}, c.config.AuthKey)
}

func (c *Client) DeleteGroup(name string) (*Group, error) {
	return deleteRequest[Group](fmt.Sprintf("%s/group/%s", c.config.RestIp, name), c.config.AuthKey)
}

func (c *Client) GetGroup(group string) (*Group, error) {
	return getRequest[Group](fmt.Sprintf("%s/group/%s", c.config.RestIp, group), c.config.AuthKey)
}

// GetGroupNodes Returns all of a group's Nodes
func (c *Client) GetGroupNodes(groupName string) (*[]Node, error) {
	return getRequest[[]Node](fmt.Sprintf("%s/group/%s/nodes", c.config.RestIp, groupName), c.config.AuthKey)
}

// AddUserNode Adds a Node to a group, then returns the new array of nodes
func (c *Client) AddGroupNode(groupName string, node NewNode) (*[]Node, error) { // TODO: NodeMergeStrategy
	return postRequestBody[[]Node](fmt.Sprintf("%s/group/%s/nodes", c.config.RestIp, groupName), node, c.config.AuthKey)
}

// AddGroupNodes Adds multiple Nodes to a group, then returns the new array of nodes
func (c *Client) AddGroupNodes(groupName string, nodes []NewNode) (*[]Node, error) { // TODO: NodeMergeStrategy
	return patchRequestBody[[]Node](fmt.Sprintf("%s/group/%s/nodes", c.config.RestIp, groupName), nodes, c.config.AuthKey)
}

// SetGroupNodes Replaces all the Nodes of the group with newNodes, then returns the new array of nodes
func (c *Client) SetGroupNodes(groupName string, newNodes []NewNode) (*[]Node, error) { // TODO: NodeMergeStrategy
	return putRequestBody[[]Node](fmt.Sprintf("%s/group/%s/nodes", c.config.RestIp, groupName), newNodes, c.config.AuthKey)
}

// RemoveGroupNodes Removes multiple Nodes from a group, then returns the new array of nodes
func (c *Client) RemoveGroupNodes(groupName string, nodes []NewNode) (*[]Node, error) { // TODO: NodeMergeStrategy
	return deleteRequestBody[[]Node](fmt.Sprintf("%s/group/%s/nodes", c.config.RestIp, groupName), nodes, c.config.AuthKey)
}

func (c *Client) GroupHasPermission(group string, permission string) (*PermissionCheckResult, error) {
	return getRequest[PermissionCheckResult](fmt.Sprintf("%s/group/%s/permission-check?permission=%s", c.config.RestIp, group, permission), c.config.AuthKey)
}

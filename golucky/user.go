package golucky

import "fmt"

type userPair struct {
	UniqueId string `json:"uniqueId"`
	Username string `json:"username"`
}

type UserLookupResult userPair

type UserResult struct {
	UniqueId     string   `json:"uniqueId"`
	Username     string   `json:"username"`
	ParentGroups []string `json:"parentGroups"`
	Nodes        []Node   `json:"nodes"`
	Metadata     Metadata `json:"metadata"`
}

type TrackResult struct {
	Success   bool   `json:"success"`
	Status    string `json:"status"` // Enum: [ success, added_to_first_group, malformed_track, end_of_track, ambiguous_call, undefined_failure ]
	GroupFrom string `json:"groupFrom"`
	GroupTo   string `json:"groupTo"`
}

type trackBody struct {
	Track string `json:"track"`
}

type CreateUserResult struct {
	Outcomes         []string `json:"outcomes"`
	PreviousUsername string   `json:"previousUsername"`
	OtherUniqueIds   []string `json:"otherUniqueIds"`
}

func (c *Client) ListUsers() (*[]string, error) {
	return getRequest[[]string](fmt.Sprintf("%s/user", c.config.RestIp), c.config.AuthKey)
}

func (c *Client) CreateUser(uuid string, username string) (*CreateUserResult, error) {
	return postRequestBody[CreateUserResult](fmt.Sprintf("%s/user", c.config.RestIp), userPair{UniqueId: uuid, Username: username}, c.config.AuthKey)
}

func (c *Client) DeleteUser(uuid string) error {
	return deleteRequestNoResponse(fmt.Sprintf("%s/user/%s", c.config.RestIp, uuid), c.config.AuthKey)
}

func (c *Client) LookupUserUUID(uuid string) (*UserLookupResult, error) {
	return getRequest[UserLookupResult](fmt.Sprintf("%s/user/lookup?uniqueId=%s", c.config.RestIp, uuid), c.config.AuthKey)
}

func (c *Client) LookupUserUsername(username string) (*UserLookupResult, error) {
	return getRequest[UserLookupResult](fmt.Sprintf("%s/user/lookup?username=%s", c.config.RestIp, username), c.config.AuthKey)
}

func (c *Client) GetUserData(uuid string) (*UserResult, error) {
	return getRequest[UserResult](fmt.Sprintf("%s/user/%s", c.config.RestIp, uuid), c.config.AuthKey)
}

// GetUserNodes Returns all of a user's Nodes
func (c *Client) GetUserNodes(uuid string) (*[]Node, error) {
	return getRequest[[]Node](fmt.Sprintf("%s/user/%s/nodes", c.config.RestIp, uuid), c.config.AuthKey)
}

// AddUserNode Adds a Node to a user, then returns the new array of nodes
func (c *Client) AddUserNode(uuid string, node NewNode) (*[]Node, error) { // TODO: NodeMergeStrategy
	return postRequestBody[[]Node](fmt.Sprintf("%s/user/%s/nodes", c.config.RestIp, uuid), node, c.config.AuthKey)
}

// AddUserNodes Adds multiple Nodes to a user, then returns the new array of nodes
func (c *Client) AddUserNodes(uuid string, nodes []NewNode) (*[]Node, error) { // TODO: NodeMergeStrategy
	return patchRequestBody[[]Node](fmt.Sprintf("%s/user/%s/nodes", c.config.RestIp, uuid), nodes, c.config.AuthKey)
}

// SetUserNodes Replaces all the Nodes of the user with newNodes, then returns the new array of nodes
func (c *Client) SetUserNodes(uuid string, newNodes []NewNode) (*[]Node, error) { // TODO: NodeMergeStrategy
	return putRequestBody[[]Node](fmt.Sprintf("%s/user/%s/nodes", c.config.RestIp, uuid), newNodes, c.config.AuthKey)
}

// RemoveUserNodes Removes multiple Nodes from a user, then returns the new array of nodes
func (c *Client) RemoveUserNodes(uuid string, nodes []NewNode) (*[]Node, error) { // TODO: NodeMergeStrategy
	return deleteRequestBody[[]Node](fmt.Sprintf("%s/user/%s/nodes", c.config.RestIp, uuid), nodes, c.config.AuthKey)
}

func (c *Client) UserHasPermission(uuid string, permission string) (*PermissionCheckResult, error) {
	return getRequest[PermissionCheckResult](fmt.Sprintf("%s/user/%s/permission-check?permission=%s", c.config.RestIp, uuid, permission), c.config.AuthKey)
}

func (c *Client) Promote(uuid string, track string) (*TrackResult, error) {
	return postRequestBody[TrackResult](fmt.Sprintf("%s/user/%s/promote", c.config.RestIp, uuid), trackBody{Track: track}, c.config.AuthKey)
}

func (c *Client) Demote(uuid string, track string) (*TrackResult, error) {
	return postRequestBody[TrackResult](fmt.Sprintf("%s/user/%s/demote", c.config.RestIp, uuid), trackBody{Track: track}, c.config.AuthKey)
}

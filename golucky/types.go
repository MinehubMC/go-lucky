package golucky

type Node struct {
	Key     string    `json:"key"`
	Value   bool      `json:"value"`
	Type    string    `json:"type"` // Enum: [ permission, regex_permission, inheritance, prefix, suffix, meta, weight, display_name ]
	Expiry  int       `json:"expiry"`
	Context []Context `json:"context"`
}

// NewNode Only Key is required. Note: Value defaults to false
type NewNode struct {
	Key     string    `json:"key"`
	Value   bool      `json:"value"`
	Expiry  int       `json:"expiry"`
	Context []Context `json:"context"`
}

type Metadata struct {
	Prefix       string `json:"prefix"`
	Suffix       string `json:"suffix"`
	PrimaryGroup string `json:"primaryGroup"`
}

type Context struct {
	Key   string `json:"key"`
	Value bool   `json:"value"`
}

//type PermissionCheckRequest struct { // TODO: do
//	Permission string `json:"permission"`
//	Node       Node   `json:"node"`
//}

type PermissionCheckResult struct {
	Result string `json:"result"` // Tristate: [ true, false, undefined ]
	Node   Node   `json:"node"`
}

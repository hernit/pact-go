package types

// Message pact type
type Message struct {
	Description string      `json:"description"`
	Content     interface{} `json:"content"`
}

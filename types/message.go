package types

// Message is the PROVIDER side message type
// TODO: rationalise this for consumer/provider
type Message struct {
	// Message Body
	Content map[string]interface{} `json:"content,omitempty"`

	// Provider state to be written into the Pact file
	State string `json:"state,omitempty"`

	// Message metadata
	Metadata map[string]string `json:"metadata,omitempty"`

	// Description to be written into the Pact file
	Description string `json:"description"`
}

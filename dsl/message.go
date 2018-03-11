package dsl

import "github.com/pact-foundation/pact-go/types"

// Message is the main implementation of the Pact Message interface.
type Message struct {
	message types.Message
}

// Given specifies a provider state. Optional.
func (p *Message) Given(state string) *Message {
	p.message.State = state
	return p
}

// UponReceiving specifies the name of the test case. This becomes the name of
// the consumer/provider pair in the Pact file. Mandatory.
func (m *Message) ExpectsToReceive(description string) *Message {
	m.message.Description = description
	return m
}

// WithMetadata specifies message-implementation specific metadata
// to go with the content
func (p *Message) WithMetadata(metadata map[string]string) *Message {
	p.message.Metadata = metadata
	return p
}

// WithRequest specifies the details of the HTTP request that will be used to
// confirm that the Provider provides an API listening on the given interface.
// Mandatory.
func (p *Message) WithContent(content interface{}) *Message {

	// Need to fix any weird JSON marshalling issues with the body Here
	// If body is a string, not an object, we need to put it back into an object
	// so that it's not double encoded
	switch body := content.(type) {
	case string:
		p.message.Content = toObject([]byte(body))
	default:
		// leave alone
		p.message.Content = content
	}

	return p
}

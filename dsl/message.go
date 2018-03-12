package dsl

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/pact-foundation/pact-go/types"
)

// Message is the main implementation of the Pact Message interface.
type Message struct {
	message types.Message
	Args    []string
}

// Given specifies a provider state. Optional.
func (p *Message) Given(state string) *Message {
	p.message.State = state
	return p
}

// ExpectsToReceive specifies the content it is expecting to be
// given from the Provider. The function must be able to handle this
// message for the interaction to succeed.
func (p *Message) ExpectsToReceive(description string) *Message {
	p.message.Description = description
	return p
}

// WithMetadata specifies message-implementation specific metadata
// to go with the content
func (p *Message) WithMetadata(metadata map[string]string) *Message {
	p.message.Metadata = metadata
	return p
}

// WithContent specifies the details of the HTTP request that will be used to
// confirm that the Provider provides an API listening on the given interface.
// Mandatory.
func (p *Message) WithContent(content interface{}) *Message {

	// Need to fix any weird JSON marshalling issues with the body Here
	// If body is a string, not an object, we need to put it back into an object
	// so that it's not double encoded
	switch body := content.(type) {
	case string:
		p.message.Content, _ = toMappedObject([]byte(body))
	case map[string]interface{}:
		p.message.Content = body
	default:
		// TODO: fail??
		// p.message.Content, _ = toMappedObject([]byte(body))
	}

	return p
}

func toMappedObject(content []byte) (map[string]interface{}, error) {
	var obj map[string]interface{}
	err := json.Unmarshal(content, &obj)
	if err != nil {
		log.Println("[DEBUG] interaction: error unmarshaling object into string:", err.Error())
		return nil, errors.New("unable to marshal content into object")
	}

	return obj, nil
}

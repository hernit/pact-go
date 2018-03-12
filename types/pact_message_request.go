package types

import "encoding/json"

// PactMessageRequest contains the response from the Pact Message
// CLI execution.
type PactMessageRequest struct {
	Message       Message
	Consumer      string
	Provider      string
	PactDir       string
	PactWriteMode string
	Args          []string
}

// Validate checks all things are well and constructs
// the CLI args to the message service
func (m *PactMessageRequest) Validate() error {
	m.Args = []string{}

	body, err := json.Marshal(m.Message)
	if err != nil {
		return err
	}

	m.Args = append(m.Args, []string{
		m.PactWriteMode,
		string(body),
		"--consumer",
		m.Consumer,
		"--provider",
		m.Provider,
		"--pact-dir",
		m.PactDir,
	}...)

	return nil
}

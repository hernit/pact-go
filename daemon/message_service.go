package daemon

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/kardianos/osext"
)

// MessageService is a wrapper for the Pact Message Service.
type MessageService struct {
	ServiceManager
}

// NewService creates a new MessageService with default settings.
func (m *MessageService) NewService(args []string) Service {
	log.Printf("[DEBUG] starting verification service with args: %v\n", args)

	m.Args = []string{
		"exec",
		"pact-message",
	}
	m.Args = append(m.Args, args...)
	// m.Args = []string{"exec", "pact-message", "update", `{ "description": "a test mesage", "content": { "name": "Mary" } }`, "--consumer", "from", "--provider", "golang", "--pact-dir", "/tmp"}

	log.Printf("[DEBUG] starting verification service with args: %v\n", m.Args)
	m.Cmd = getVerifierCommandPath()
	return m
}

// Runs bundler in the ./bin directory
func getMessageCommandPath() string {
	dir, _ := osext.ExecutableFolder()
	return fmt.Sprintf(filepath.Join(dir, "bin", "bundle"))
}

package provider

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

var commonHeaders = map[string]string{
	"Content-Type": "application/json; charset=utf-8",
}

var pact = createPact()

func TestMessageConsumer_Success(t *testing.T) {
	message := &dsl.Message{}
	message.
		Given("some state").
		ExpectsToReceive("some test case").
		WithMetadata(commonHeaders).
		WithContent(map[string]interface{}{
			"foo": "bar",
		})

	res, err := pact.VerifyMessage(message, func(i ...types.Message) error {
		t.Logf("[DEBUG] calling message handler func with arguments: %v \n", i)

		return nil
	})

	if err != nil {
		t.Fatal("VerifyMessage failed:", err)
	}

	t.Log("[DEBUG] Response from VerifyMessage:", res)

}
func TestMessageConsumer_Fail(t *testing.T) {
	t.Skip()
	message := &dsl.Message{}
	message.
		Given("some state").
		ExpectsToReceive("some test case").
		WithMetadata(commonHeaders).
		WithContent(map[string]interface{}{
			"foo": "bar",
		})

	res, err := pact.VerifyMessage(message, func(i ...types.Message) error {
		t.Logf("[DEBUG] calling message handler func with arguments: %v \n", i)

		return errors.New("something bad happened and I couldn't parse the message")
	})

	if err != nil {
		t.Fatal("VerifyMessage failed:", err)
	}

	t.Log("[DEBUG] Response from VerifyMessage:", res)

}

// Configuration / Test Data
// var port, _ = utils.GetFreePort()
var port = 9393
var dir, _ = os.Getwd()
var pactDir = fmt.Sprintf("%s/../../pacts", dir)
var logDir = fmt.Sprintf("%s/log", dir)

// Setup the Pact client.
func createPact() dsl.Pact {
	// Create Pact connecting to local Daemon
	return dsl.Pact{
		Port:     6666,
		Consumer: "billy",
		Provider: "bobby",
		LogDir:   logDir,
		// PactDir:  pactDir, // TODO: this seems to cause an issue "NoMethodError: undefined method `content' for #<Pact::Interaction:0x00007fc8f1a082e8>"
		PactDir:           "/tmp",
		LogLevel:          "DEBUG",
		PactFileWriteMode: "update",
	}
}

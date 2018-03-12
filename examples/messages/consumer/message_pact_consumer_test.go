package provider

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

// Common test data
var dir, _ = os.Getwd()
var pactDir = fmt.Sprintf("%s/../../pacts", dir)
var logDir = fmt.Sprintf("%s/log", dir)
var pact dsl.Pact
var form url.Values
var rr http.ResponseWriter
var req *http.Request

var name = "billy"
var like = dsl.Like
var eachLike = dsl.EachLike
var term = dsl.Term
var loginRequest = fmt.Sprintf(`{ "username":"%s", "password": "issilly" }`, name)

var commonHeaders = map[string]string{
	"Content-Type": "application/json; charset=utf-8",
}

// Use this to control the setup and teardown of Pact
func TestMain(m *testing.M) {
	// Setup Pact and related test stuff
	setup()

	// Run all the tests
	code := m.Run()

	os.Exit(code)
}

// Setup common test data
func setup() {
	pact = createPact()

	// Record response (satisfies http.ResponseWriter)
	rr = httptest.NewRecorder()
}

func TestMessageConsumer(t *testing.T) {
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

// Configuration / Test Data
// var port, _ = utils.GetFreePort()
var port = 9393

// Setup the Pact client.
func createPact() dsl.Pact {
	// Create Pact connecting to local Daemon
	return dsl.Pact{
		Port:     6666,
		Consumer: "billy",
		Provider: "bobby",
		LogDir:   logDir,
		PactDir:  pactDir, // TODO: this seems to cause an issue "NoMethodError: undefined method `content' for #<Pact::Interaction:0x00007fc8f1a082e8>"
		// PactDir:           "/tmp",
		LogLevel:          "DEBUG",
		PactFileWriteMode: "update",
	}
}

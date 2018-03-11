package provider

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
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
		WithMetadata(map[string]string{
			"content-type": "application/json",
		}).
		// WithContent(map[string]interface{}{
		// 	"foo": "bar",
		// })
		WithContent(`{"s":"foo"}`)

	res, err := pact.VerifyMessage(message, func(i ...interface{}) error {
		t.Log("[DEBUG] calling message handler func")
		t.Log("[DEBUG] arguments:", i)

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
		PactDir:  pactDir,
		LogLevel: "DEBUG",
	}
}

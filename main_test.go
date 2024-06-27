package main

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io"
	"lucas.blin/v2/config"
	"net/http"
	"testing"
)

func TestMainWithArgs(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Register responder for our fake URL
	responder := httpmock.NewStringResponder(200, `
	{
		"message-count": "1",
		"messages": [
			{
				"to": "33616161616",
				"message-id": "50414951-13d0-4803-b4ed-ae70cddc2605",
				"status": "0",
				"remaining-balance": "20.73884000",
				"message-price": "0.07119000",
				"network": "20810"
			}
		]
	}`)
	httpmock.RegisterResponder("POST", "http://test:8080", responder)

	config.Url = "http://test:8080"
	argv := &ArgT{
		Sender:     "33616161616",
		Recipient:  "33616161622",
		Message:    "'Testing message'",
		ConfigPath: "~/.config/parameters-test.ini",
	}

	_ = RunCli(argv)

	// Create a *http.Request
	req, err := http.NewRequest("POST", "http://test:8080", nil)
	if err != nil {
		t.Fatal(err)
	}

	err, p := CreatePayload(argv.Sender, argv.Recipient, argv.Message, argv.ConfigPath)
	if err != nil {
		return
	}

	resp, err := responder(req)
	if err != nil {
		t.Fatal(err)
	}

	// Read the body of the *http.Response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Assert the body
	expectedBody := `
	{
		"message-count": "1",
		"messages": [
			{
				"to": "33616161616",
				"message-id": "50414951-13d0-4803-b4ed-ae70cddc2605",
				"status": "0",
				"remaining-balance": "20.73884000",
				"message-price": "0.07119000",
				"network": "20810"
			}
		]
	}`

	expectedPayload := `{"from":"33616161616","text":"'Testing message'","to":"33616161622","api_key":"api-key-value","api_secret":"api-secret-value"}`

	assert.Equal(t, expectedBody, string(body))
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, expectedPayload, string(p))
}

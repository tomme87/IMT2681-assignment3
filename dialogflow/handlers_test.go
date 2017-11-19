package dialogflow

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleDialogflowRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(HandleDialogflowRequest))
	defer ts.Close()

	dfString := []byte(`{
  "id": "f4d383b5-e9c7-4d41-86a9-84fcf434cf30",
  "timestamp": "2017-11-18T14:51:05.188Z",
  "lang": "en",
  "result": {
    "source": "agent",
    "resolvedQuery": "What is the exchange rate between SEK and NOK",
    "action": "",
    "actionIncomplete": false,
    "parameters": {
      "baseCurrency": "SEK",
      "targetCurrency": "NOK"
    },
    "contexts": [],
    "metadata": {
      "intentId": "19d48d37-9ab9-4d64-b4d2-535d705f2c28",
      "webhookUsed": "true",
      "webhookForSlotFillingUsed": "false",
      "webhookResponseTime": 686,
      "intentName": "exchange"
    },
    "fulfillment": {
      "speech": "Latest rate between SEK and NOK is 0.9770723",
      "source": "exchange-rate-api",
      "displayText": "Latest rate between SEK and NOK is 0.9770723",
      "messages": [
        {
          "type": 0,
          "speech": "Latest rate between SEK and NOK is 0.9770723"
        }
      ],
      "data": {
        "slack": {
          "text": "Latest rate between SEK and NOK is 0.9770723"
        }
      }
    },
    "score": 1
  },
  "status": {
    "code": 200,
    "errorType": "success",
    "webhookTimedOut": false
  },
  "sessionId": "0600a4b3-d924-4315-bc5f-dd7c90b4e222"
}`)

	resp, err := http.Post(ts.URL, "application/json", bytes.NewBuffer(dfString))
	if err != nil {
		t.Errorf("Error executing request: %s", err.Error())
		return
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		t.Errorf("Expected %d, but got %d, Body: %s", http.StatusOK, resp.StatusCode, body)
	}

	dfr := DialogFlowResponse{}
	err = json.NewDecoder(resp.Body).Decode(&dfr)
	if err != nil {
		body, _ := ioutil.ReadAll(resp.Body)
		t.Errorf("Error executing decoding json. Error %s. Body: %s", err, body)
		return
	}

	if !strings.HasPrefix(dfr.Speech, "Latest rate between SEK and NOK is ") {
		t.Errorf("Wrong response, got: %s", dfr.Speech)
	}
}

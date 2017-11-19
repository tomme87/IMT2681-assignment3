curl -X POST -H "Content-Type: application/json" localhost:8081/dialogflow -d '
{
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
      "speech": "Latest rate between SEK and NOK is 0.9770723\n",
      "source": "exchange-rate-api",
      "displayText": "Latest rate between SEK and NOK is 0.9770723\n",
      "messages": [
        {
          "type": 0,
          "speech": "Latest rate between SEK and NOK is 0.9770723\n"
        }
      ],
      "data": {
        "slack": {
          "text": "Latest rate between SEK and NOK is 0.9770723\n"
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
}
'
echo

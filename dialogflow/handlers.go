package dialogflow

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func HandleDialogflowRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//body, _ := ioutil.ReadAll(r.Body)
		//fmt.Printf("Unable to decode JSON: %s", body)
		//r.Body.Close()

		df := Dialogflow{}
		err := json.NewDecoder(r.Body).Decode(&df)
		if err != nil {
			fmt.Println("Error decoding")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// TODO find exchange rate
		err = df.Validate()
		if err != nil {
			fmt.Println("Erorr validating")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		p := df.Result.Parameters

		/*
			LatestPayload struct used for storing /average and /latest request payloads
		*/
		type LatestPayload struct {
			BaseCurrency   string `json:"baseCurrency"`
			TargetCurrency string `json:"targetCurrency"`
		}
		lp := LatestPayload{
			p.BaseCurrency,
			p.TargetCurrency,
		}

		jsonStr, _ := json.Marshal(lp)

		// Get URL from env variable.
		latestURL := os.Getenv("LATEST_URL")
		if latestURL == "" {
			latestURL = "https://whispering-savannah-13838.herokuapp.com/exchange/latest"
		}

		res, err := http.Post(latestURL,
			"application/json", bytes.NewBuffer(jsonStr))
		if err != nil {
			return
		}

		rate, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()

		text := fmt.Sprintf("Latest rate between %s and %s is %s",
			p.BaseCurrency, p.TargetCurrency, string(rate))

		dfr := DialogFlowResponse{
			Speech:      text,
			DisplayText: text,
			Source:      "exchange-rate-api",
		}
		dfr.Data.Slack.Text = text

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(dfr)
	}
}

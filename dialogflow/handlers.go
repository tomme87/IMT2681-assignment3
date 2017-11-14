package dialogflow

import (
	"net/http"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"bytes"
)

func HandleDialogflowRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Printf("Unable to decode JSON: %s", body)
		r.Body.Close()

		df := Dialogflow{}
		err := json.NewDecoder(r.Body).Decode(&df)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// TODO find exchange rate
		err = df.Validate()
		if err != nil {
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

		res, err := http.Post("https://howling-skull-56836.herokuapp.com/exchange/latest",
			"application/json", bytes.NewBuffer(jsonStr))
		if err != nil {
			return
		}

		rate, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()


		text := fmt.Sprintf("Latest rate between %s and  %s  is %s",
			p.BaseCurrency, p.TargetCurrency, string(rate))

		dfr := DialogFlowResponse{
			Speech: "",
			DisplayText: text,
		}
		dfr.Data.Slack.Text = text


		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(dfr)
	}
}

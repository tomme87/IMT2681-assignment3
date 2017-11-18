package dialogflow

import "errors"

type Dialogflow struct {
	Result struct {
		Parameters struct {
			BaseCurrency   string
			TargetCurrency string
		}
	}
}

func (df *Dialogflow) Validate() error {
	p := df.Result.Parameters

	if p.TargetCurrency == "" || p.BaseCurrency == "" {
		return errors.New("Currency is not precent in request")
	}

	return nil
}

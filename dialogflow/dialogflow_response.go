package dialogflow

type DialogFlowResponse struct {
	Speech      string `json:"speech"`
	DisplayText string `json:"displayText"`
	Data        struct {
		Slack struct {
			Text string `json:"text"`
		} `json:"slack"`
	} `json:"data"`
	//ContextOut []string
	Source string `json:"source"`
	//FollowupEvent string
}

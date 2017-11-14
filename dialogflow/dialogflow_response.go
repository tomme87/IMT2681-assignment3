package dialogflow

type DialogFlowResponse struct {
	Speech string
	DisplayText string
	Data struct{
		Slack struct{
			Text string
		}
	}
	//ContextOut []string
	Source string
	//FollowupEvent string
}

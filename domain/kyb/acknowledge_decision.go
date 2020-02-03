package kyb

type AcknowledgeDecisionRequest struct{
	Uuid string `json:"uuid"`
}

type AcknowledgeDecisionResponse struct{
	Result uint64 `json:"result"`
}
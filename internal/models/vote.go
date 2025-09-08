package models

type Vote struct {
	ParticipantID string `json:"participant_id"`
	Value         string `json:"value"`
}

type VoteResponse struct {
	ParticipantID   string `json:"participant_id"`
	ParticipantName string `json:"participant_name"`
	Value           string `json:"value"`
}

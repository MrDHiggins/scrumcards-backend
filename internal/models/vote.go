package models

type Vote struct {
	ParticipantID string `json:"participant_id"`
	Value         string `json:"value"`
}

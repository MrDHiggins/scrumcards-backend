package models

import "time"

type Session struct {
	ID           string                  `json:"id"`
	CreatedAt    time.Time               `json:"created_at"`
	HostId       string                  `json:"host_id"`
	Ticket       string                  `json:"ticket,omitempty"`
	Participants map[string]*Participant `json:"participants,omitempty"`
	Votes        map[string]*Vote        `json:"votes,omitempty"`
	Revealed     bool                    `json:"revealed"`
}

type SessionResponse struct {
	ID           string         `json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	HostID       string         `json:"host_id"`
	Ticket       string         `json:"ticket,omitempty"`
	Participants []*Participant `json:"participants,omitempty"`
	Votes        []*Vote        `json:"votes,omitempty"`
	Revealed     bool           `json:"revealed"`
}

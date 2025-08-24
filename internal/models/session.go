package models

import "time"

type Session struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	HostId    string    `json:"host_id"`
	Ticket    string    `json:"ticket,omitempty"` //Omitted for MVP
}

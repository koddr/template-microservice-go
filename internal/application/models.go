package application

import "time"

// Transaction struct.
type Transaction struct {
	ID            int       `json:"id"`
	EventID       int       `json:"event_id"`
	CreatedAt     time.Time `json:"created_at"`
	ProfileID     string    `json:"profile_id"`
	MessengerID   string    `json:"messenger_id"`
	MessengerName string    `json:"messenger_name"`
	EventType     string    `json:"event_type"`
	UTMSource     string    `json:"utm_source"`
	UTMMedium     string    `json:"utm_medium"`
	UTMCampaign   string    `json:"utm_campaign"`
	UTMContent    string    `json:"utm_content"`
	UTMTerm       string    `json:"utm_term"`
}

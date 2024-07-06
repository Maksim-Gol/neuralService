package neuralservice

import "time"

type ServiceCall struct {
	UserID    string         `json:"user_id"`
	ModelID   string         `json:"model_id"`
	RequestID string         `json:"request_id"`
	Cost      int            `json:"cost"`
	Status    string         `json:"status"`
	Metadata  map[string]any `json:"metadata"`
	CallTime  time.Time      `json:"call_time"`
}

package guessers

import "time"

type NationGuesserRequest struct {
	RequestId      string `json:"request_id"`
	ConversationId string `json:"conversation_id"`
	Code           string `json:"code,omitempty"`
	Message        string `json:"message,omitempty"`
}

type NationGuesserResult struct {
	Nation string  `json:"nation"`
	Value  float64 `json:"value"`
}

type NationGuesserResponse struct {
	RequestId      string    `json:"request_id"`
	Date           time.Time `json:"date"`
	Answer         string    `json:"answer"`
	ConversationId string    `json:"conversation_id"`
	MessageId      string    `json:"message_id"`
}

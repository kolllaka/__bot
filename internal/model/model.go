package model

import "time"

type Message struct {
	ID        uint      `json:"id,omitempty"`
	Platform  string    `json:"platform,omitempty"`
	Channel   string    `json:"channel,omitempty"`
	MsgAuthor string    `json:"msg_author,omitempty"`
	Msg       string    `json:"msg,omitempty"`
	Date      time.Time `json:"date,omitempty"`
	IsCommand bool      `json:"is_command,omitempty"`
	NameCmd   string    `json:"name_cmd,omitempty"`
	Error     error     `json:"error,omitempty"`
}

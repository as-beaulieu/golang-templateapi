package models

import "time"

type HeartbeatResponse struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

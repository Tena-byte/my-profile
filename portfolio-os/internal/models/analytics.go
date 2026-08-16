package models

import "time"

type Visit struct {
	ID             string        `json:"id"`
	IPAddress      string        `json:"ip_address"`
	UserAgent      string        `json:"user_agent"`
	DeviceCategory string        `json:"device_category"`
	Page           string        `json:"page"`
	Timestamp      time.Time     `json:"timestamp"`
	Duration       time.Duration `json:"duration"`
}

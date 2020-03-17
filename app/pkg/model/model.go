package model

import (
	"time"
)

type VPNdata struct {
	Timestamp     time.Time `json:"@timestamp"`
	Host          string    `json:"host"`
	SyslogProgram string    `json:"syslog_program"`
	Message       string    `json:"message"`
	Username      string    `json:"username"`
	SrcIP         string    `json:"src_ip"`
	Location      uint      `json:"location"`
	VpnHash       string    `json:"vpn_hash"`
}

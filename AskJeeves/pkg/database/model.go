package database

import (
	"time"
)

type UserVPNLog struct {
	CreatedAt        time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt        *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Username         string     `json:"username" `
	VpnHash          string     `gorm:"primary_key" json:"vpn_hash"`
	IPaddr           string     `json:"src_ip"`
	Location         string
	ISOcode          uint
	UserConfirmation bool `gorm:"default:false"`
	CaseID           int
	Counter          int
	EventID          string
	Device           string `json:"device"`
	Hostname         string `json:"hostname"`
}

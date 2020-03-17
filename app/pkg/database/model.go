package database

import (
	"github.com/jinzhu/gorm"
)

type UserVPNLog struct {
	gorm.Model
	Username      string
	VpnHash       string `gorm:"primary_key"`
	IPaddr        string
	Location      uint
	Confirm       bool
	Count         int
	LastLoginDate string
}

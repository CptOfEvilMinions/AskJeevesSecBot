package database

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/CptOfEvilMinions/AskJeevesSecBot/pkg/config"
)

// InitMySQLConnector input:
// InitMySQLConnector output:
// https://tutorialedge.net/golang/golang-mysql-tutorial/
func InitMySQLConnector(cfg *config.Config) (*gorm.DB, error) {
	// Create MySQL DSN string
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	fmt.Println("=============== MySQL ===============")
	fmt.Println(cfg.MySQL.Username)
	fmt.Println(cfg.MySQL.Password)
	fmt.Println(cfg.MySQL.Protocol)
	fmt.Println(cfg.MySQL.Hostname)
	fmt.Println(strconv.Itoa(cfg.MySQL.Port))
	fmt.Println(cfg.MySQL.Database)

	MySQLServerDSN := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.MySQL.Username,
		cfg.MySQL.Password,
		cfg.MySQL.Hostname,
		strconv.Itoa(cfg.MySQL.Port),
		cfg.MySQL.Database,
	)
	fmt.Println(MySQLServerDSN)

	// Connect to database
	db, err := gorm.Open("mysql", MySQLServerDSN)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	// defer the close till after the main function has finished executing
	//defer db.Close()

	// Init schema
	db.AutoMigrate(&UserVPNLog{})

	return db, nil
}

// AddVpnUserEntry input:
// AddVpnUserEntry output:
func AddVpnUserEntry(db *gorm.DB, Username string, VpnHash string, IPaddr string, Location uint) (bool, error) {
	// Create user VPN entry
	userVPNLog := UserVPNLog{
		Username:      Username,
		VpnHash:       VpnHash,
		IPaddr:        IPaddr,
		Location:      Location,
		Confirm:       false,
		Count:         1,
		LastLoginDate: time.Now().Format(time.RFC3339),
	}

	// Add record to database
	err := db.Create(&userVPNLog).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// updateUserVPNCounter input:
// updateUserVPNCounter output:
func updateUserVPNCounter(db *gorm.DB, userVPNLog UserVPNLog) {
	db.First(&userVPNLog)
	userVPNLog.Count++
	db.Save(&userVPNLog)
}

// QueryDoesVpnHashExist input: DB connecter and VPnHash
// QueryDoesVpnHashExist output: If Vpnhash exists return true, else false
func QueryDoesVpnHashExist(db *gorm.DB, VpnHash string) bool {
	// init obj
	userVPNLog := UserVPNLog{}

	// See if VpnHash exists
	if db.Where("vpn_hash = ?", VpnHash).First(&userVPNLog).RecordNotFound() {
		return false
	}
	updateUserVPNCounter(db, userVPNLog)
	return true
}

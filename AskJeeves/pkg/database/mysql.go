package database

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/CptOfEvilMinions/AskJeevesSecBot/pkg/config"
)

// InitMySQLConnector input: config
// InitMySQLConnector output: Return MySQL connector
// https://tutorialedge.net/golang/golang-mysql-tutorial/
func InitMySQLConnector(cfg *config.Config) (*gorm.DB, error) {
	// Create MySQL DSN string
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	fmt.Println("=============== MySQL ===============")
	fmt.Println("Username:", cfg.MySQL.Username)
	fmt.Println("Password:", cfg.MySQL.Password)
	fmt.Println("Protocol:", cfg.MySQL.Protocol)
	fmt.Println("Hostname:", cfg.MySQL.Hostname)
	fmt.Println("Port:", strconv.Itoa(cfg.MySQL.Port))
	fmt.Println("Database:", cfg.MySQL.Database)

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

	// Init schema
	db.AutoMigrate(&UserVPNLog{})

	return db, nil
}

// AddVpnUserEntry input: MySQL connector and values about the VPN login
// AddVpnUserEntry output: Return True is login is sucessfully added to DB
func AddVpnUserEntry(db *gorm.DB, Username string, VpnHash string, IPaddr string, ISOcode uint, Location string) (bool, error) {
	// Create user VPN entry
	userVPNLog := UserVPNLog{
		Username: Username,
		VpnHash:  VpnHash,
		IPaddr:   IPaddr,
		ISOcode:  ISOcode,
		Location: Location,
		Confirm:  false,
		Counter:  1,
	}

	// Add record to database
	err := db.Create(&userVPNLog).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// DeleteOldEntries input: MySQL connector and config
// This function will init a Golang Ticker to perform this task
// every cfg.MySQL.Interval (seconds) to query all entries in the
// database and delete all entries older than cfg.MySQL.Expire (days)
// DeleteOldEntries output: None
// https://tutorialedge.net/golang/go-ticker-tutorial/
func DeleteOldEntries(db *gorm.DB, cfg *config.Config) {

	// Create ticker
	ticker := time.NewTicker(time.Duration(cfg.MySQL.Interval) * time.Second)

	for range ticker.C {
		currentDate := time.Now()     // Get current date YYYY-MM-DD
		userVPNLogs := []UserVPNLog{} // Init list for objs

		// Get all records
		db.Find(&userVPNLogs)
		for _, userVPNLog := range userVPNLogs {
			// Calculate Delta between timestamps
			daysDelta := currentDate.Sub(userVPNLog.UpdatedAt).Hours() / 24

			// If great than setting delete
			if daysDelta >= float64(cfg.MySQL.Expire) {
				log.Println("Deleted:", userVPNLog.VpnHash, userVPNLog.Username, userVPNLog.IPaddr, userVPNLog.ISOcode, userVPNLog.Location)
				db.Unscoped().Delete(&userVPNLog)
			}
		}

		fmt.Printf("[+] - %s - Cleaned up old entries", time.Now().Format("2006-01-02 15:04:05"))

	}

}

// updateUserVPNCounter input: MySQL connector and userVPNlog obj
// This function will update a user login counter
// updateUserVPNCounter output:  None
func updateUserVPNCounter(db *gorm.DB, userVPNLog UserVPNLog) {
	db.First(&userVPNLog)
	userVPNLog.Counter++
	db.Save(&userVPNLog)
}

// QueryDoesVpnHashExist input: MySQL connector and VpnHash
// QueryDoesVpnHashExist output: If the VpnHash exists return true, else false
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

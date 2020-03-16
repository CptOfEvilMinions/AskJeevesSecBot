package database

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/CptOfEvilMinions/AskJeevesSecBot/pkg/config"
)

type UserVPNLog struct {
	gorm.Model
	Username      string
	VpnHash       string
	IPaddr        string
	Location      string
	Confirm       bool
	Count         int
	LastLoginDate string
}

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
	defer db.Close()

	// Init schema
	db.AutoMigrate(&UserVPNLog{})

	return db, nil
}

// Query input:
// Query ouutput:
// func Query(db *sql.DB, query) (string, error) {
// 	results, err := db.Query(query)
// 	if err != nil {
// 		panic(err.Error()) // proper error handling instead of panic in your app
// 		return "", err
// 	}
// 	return results, nil
// }

// QueryDoesVpnHashExist input:
// QueryDoesVpnHashExist output:
func QueryDoesVpnHashExist(db *gorm.DB, VpnHash string) bool {
	// init obj
	var userVPNLog UserVPNLog

	// Check if returns RecordNotFound error
	db.Where("VpnHash = ?", VpnHash).First(&userVPNLog).RecordNotFound()
	if db.Model(&userVPNLog).Related(&VpnHash).RecordNotFound() {
		return false, nil
	}

	return true, nil
}

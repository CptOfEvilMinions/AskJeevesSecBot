package butlingbutler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/CptOfEvilMinions/AskJeevesSecBot/pkg/config"
)

type UserResponse struct {
}

// InitJWTtoken input:
// InitJWTtoken output:
func InitJWTtoken(cfg *config.Config) (string, error) {
	// Generate JSON payload
	requestBody, err := json.Marshal(map[string]string{
		"username": cfg.ButlingButler.Username,
		"password": cfg.ButlingButler.Password,
	})
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	resp, err := http.Post(cfg.ButlingButler.URL+"/auth/login", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	accessToken := result["access_token"].(string)

	return accessToken, nil
}

// getUserRespones input:
// getUserRespones output:
func getUserRespones(JWTtoken string, cfg *config.Config) (map[string]interface{}, error) {
	// Create URL
	url := cfg.ButlingButler.URL + "/askjeeves/GetUserResponse"

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + JWTtoken

	h := http.Client{Timeout: time.Second * 2} // Maximum of 2 secs

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	res, getErr := h.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	fmt.Println(body)

	var parsed map[string]interface{}
	//data := []byte
	if err := json.Unmarshal(body, &parsed); err != nil {
		log.Fatal(err)
	}

	fmt.Println(parsed["VPNHash"].(string))

	// for _, parse := range parsed {
	// 	fmt.Println(parse["VPNHash"].(string))
	// }

	return parsed, nil

}

// UpdateDatabaseEntries input:
// UpdateDatabaseEntries output:
func UpdateDatabaseEntries(JWTtoken string, cfg *config.Config) {
	// Create ticker
	ticker := time.NewTicker(time.Duration(cfg.ButlingButler.Interval) * time.Second)

	for range ticker.C {
		// Get all the user responses
		userResponses, err := getUserRespones(JWTtoken, cfg)
		if err != nil {
			log.Fatalln(err)
		}

		if len(userResponses) == 0 {
			continue
		}

		println(userResponses)

		// for _,  userResponse := range userResponses {
		// 	// Query entry

		// 	// Update entry

		// 	// Commit entry
		// }

		// currentDate := time.Now()     // Get current date YYYY-MM-DD
		// userVPNLogs := []UserVPNLog{} // Init list for objs

		// // Get all records
		// db.Find(&userVPNLogs)
		// for _, userVPNLog := range userVPNLogs {
		// 	// Calculate Delta between timestamps
		// 	daysDelta := currentDate.Sub(userVPNLog.UpdatedAt).Hours() / 24

		// 	// If great than setting delete
		// 	if daysDelta >= float64(cfg.MySQL.Expire) {
		// 		log.Println("Deleted:", userVPNLog.VpnHash, userVPNLog.Username, userVPNLog.IPaddr, userVPNLog.ISOcode, userVPNLog.Location)
		// 		db.Unscoped().Delete(&userVPNLog)
		// 	}
		// }

		fmt.Printf("[+] - %s - Got user responses", time.Now().Format("2006-01-02 15:04:05"))

	}
}

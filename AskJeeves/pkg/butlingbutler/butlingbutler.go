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
	"github.com/CptOfEvilMinions/AskJeevesSecBot/pkg/database"
	"github.com/CptOfEvilMinions/AskJeevesSecBot/pkg/ticket"
	"github.com/jinzhu/gorm"
)

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

	// HTTP POST request
	resp, err := http.Post(cfg.ButlingButler.URL+"/auth/login", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	// Extract JSON
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	accessToken := result["access_token"].(string)

	// Return accessToken
	return accessToken, nil
}

// getUserRespones input:
// getUserRespones output:
// https://gist.github.com/montanaflynn/b390b1212dada5864d9b
func getUserRespones(JWTtoken string, cfg *config.Config) ([]UserResponse, error) {
	// Create array for response
	var userResponses []UserResponse

	// Create URL
	url := cfg.ButlingButler.URL + "/askjeeves/GetUserResponse"

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + JWTtoken

	// Create an HTTP client
	c := &http.Client{}

	// Create an HTTP request
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return userResponses, err
	}

	// Add authorization header to the req
	r.Header.Add("Authorization", bearer)

	// Send the request
	res, getErr := c.Do(r)
	if getErr != nil {
		log.Fatal(getErr)
	}

	// Make sure to close after reading
	defer res.Body.Close()

	// Read all the response body
	rb, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return userResponses, err
	}

	// Read JSON into obj array
	err = json.Unmarshal(rb, &userResponses)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(getErr)
	}
	return userResponses, nil

}

// UpdateDatabaseEntries input:
// UpdateDatabaseEntries output:
func UpdateDatabaseEntries(JWTtoken string, cfg *config.Config, db *gorm.DB) {
	// Create ticker
	ticker := time.NewTicker(time.Duration(cfg.ButlingButler.Interval) * time.Second)

	for range ticker.C {
		// Get all the user responses
		userResponses, err := getUserRespones(JWTtoken, cfg)
		if err != nil {
			log.Fatalln(err)
		}

		// Continue if the user responses array has
		// a length of 0
		if len(userResponses) == 0 {
			continue
		}

		fmt.Printf("[+] - %s - Got user responses", time.Now().Format(time.RubyDate))

		for _, userResponse := range userResponses {
			fmt.Println(userResponse.VPNhash)

			// create VPN struct obj
			var userVPNLog database.UserVPNLog

			// Query entry
			//db.Where("vpn_hash = ? AND event_id = ?", userResponse.VPNhash, userResponse.EventID).First(&userVPNLog)
			db.Where("vpn_hash = ?", userResponse.VPNhash).First(&userVPNLog)

			fmt.Println(userResponse.VPNhash)
			fmt.Println(userResponse.EventID)

			// Update entry
			// True: Legit login
			// False: UNauthorized login
			if userResponse.UserSelection == "legitimate_login" {
				userVPNLog.UserConfirmation = true
			} else {
				userVPNLog.UserConfirmation = false

				// Create ticket and set caseID
				userVPNLog.CaseID, err = ticket.CreateTheHiveCase(cfg, userVPNLog)
				if err != nil {
					log.Fatalln(err)
				}
			}

			// Commit entry
			db.Save(&userVPNLog)
		}

	}
}

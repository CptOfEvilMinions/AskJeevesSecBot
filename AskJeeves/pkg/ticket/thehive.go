package ticket

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
)

// CreateTheHiveCase input:
// CreateTheHiveCase output:
func CreateTheHiveCase(cfg *config.Config, userVPNLog database.UserVPNLog) (int, error) {
	// Create URL
	url := fmt.Sprintf("%s/api/case", cfg.TheHive.URL)

	// Create bearer token
	bearerToken := fmt.Sprintf("Bearer %s", cfg.TheHive.APIkey)

	fmt.Println(userVPNLog)

	// Generate description
	description := fmt.Sprintf("**Timestamp**: %s\n\n", userVPNLog.CreatedAt.Format(time.RubyDate)) +
		fmt.Sprintf("**EventID**: %s\n\n", userVPNLog.EventID) +
		fmt.Sprintf("**Username**: %s\n\n", userVPNLog.Username) +
		fmt.Sprintf("**VPNhash**: %s\n\n", userVPNLog.VpnHash) +
		fmt.Sprintf("**Location**: %s\n\n", userVPNLog.Location) +
		fmt.Sprintf("**Device**: %s\n\n", userVPNLog.Device) +
		fmt.Sprintf("**Hostname**: %s\n\n", userVPNLog.Hostname) +
		fmt.Sprintf("**IP address**: %s\n\n", userVPNLog.IPaddr) +
		fmt.Sprintf("**References**:\n\n") +
		fmt.Sprintf("**Remediation/notes**:\n\n\n")

	// Create HTTP client
	hc := http.Client{}

	// Create POST payload
	caseTitle := fmt.Sprintf("UNauthorized VPN login - %s - %s\n", userVPNLog.Username, userVPNLog.VpnHash)
	requestBody, err := json.Marshal(map[string]string{
		"title":       caseTitle,
		"description": description,
	})

	// Make POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println(err)
	}

	// Add headers
	req.Header.Add("Authorization", bearerToken)
	req.Header.Add("Content-Type", "application/json")

	// Make ticket
	resp, err := hc.Do(req)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(resp.Status)
	fmt.Println(resp.Body)

	// resp,err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	// if err != nil {
	// 	log.Println(err)
	// }
	// defer resp.Body.Close()

	// // Make HTTP client and POST request
	// client :=&http.Client{}
	// req,_ := http.NewRequest("POST", url, nil)

	rb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, err
	}

	result := make(map[string]interface{})
	err = json.Unmarshal([]byte(rb), &result)

	caseID := result["caseId"].(float64)
	if err != nil {
		return -1, err
	}

	// Convert string to int
	//caseIDInt, err := strconv.Atoi(caseID)

	return int(caseID), nil
}

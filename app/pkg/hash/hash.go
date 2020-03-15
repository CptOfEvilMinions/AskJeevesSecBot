package hash

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// VpnHash input: Takes in username, IP address, and location (City, State, Continent)
// VpnHash output: Returns an MD5 hash of input
func VpnHash(username string, IPaddr string, location string) string {
	// String: "bob,1.1.1.1,New York, NY"
	inputString := fmt.Sprintf("%s,%s,%s", username, IPaddr, location)
	h := md5.New()
	h.Write([]byte(inputString))
	md5Hash := hex.EncodeToString(h.Sum(nil))
	return md5Hash
}

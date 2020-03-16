package hash

import (
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/sha3"
)

// VpnHash input: Takes in username, IP address, and location (City, State, GeoNameID)
// VpnHash output: Returns an SHA3 hash of input
func VpnHash(username string, IPaddr string, location uint) string {
	// String: "bob,1.1.1.1,New York, NY"
	inputString := fmt.Sprintf("%s,%s,%d", username, IPaddr, location)
	h := sha3.New256()
	h.Write([]byte(inputString))
	sha3Hash := hex.EncodeToString(h.Sum(nil))
	return sha3Hash
}

package geoip

import (
	"net"

	"github.com/oschwald/maxminddb-golang"
)

// IPaddrLocationLookup input: Takes in an IP address as a string
// IPaddrLocationLookup output: Returns city of the IP address
func IPaddrLocationLookup(IPaddr string) (string, error) {
	// Open database
	db, err := maxminddb.Open("data/GeoLite2-City.mmdb"
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Convert IP astring to net IP format
	ip := net.ParseIP(IPaddr)

	// IP lookup
	var record interface{}
	err = db.Lookup(ip, &record)
	if err == nil {
		return record, nil
	}
	return nil, err
}

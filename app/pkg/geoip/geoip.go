package geoip

import (
	"fmt"
	"net"

	"github.com/oschwald/maxminddb-golang"
)

// https://godoc.org/github.com/oschwald/geoip2-golang#City
type GeoIPCity struct {
	City struct {
		GeoNameID uint              `maxminddb:"geoname_id"`
		Names     map[string]string `maxminddb:"names"`
	} `maxminddb:"city"`
	Continent struct {
		Code      string            `maxminddb:"code"`
		GeoNameID uint              `maxminddb:"geoname_id"`
		Names     map[string]string `maxminddb:"names"`
	} `maxminddb:"continent"`
	Country struct {
		GeoNameID         uint              `maxminddb:"geoname_id"`
		IsInEuropeanUnion bool              `maxminddb:"is_in_european_union"`
		IsoCode           string            `maxminddb:"iso_code"`
		Names             map[string]string `maxminddb:"names"`
	} `maxminddb:"country"`
	Subdivisions []struct {
		GeoNameID uint              `maxminddb:"geoname_id"`
		IsoCode   string            `maxminddb:"iso_code"`
		Names     map[string]string `maxminddb:"names"`
	} `maxminddb:"subdivisions"`
}

// IPaddrLocationLookup input: Takes in an IP address as a string
// IPaddrLocationLookup output: Returns Geo lcoation UID
func IPaddrLocationLookup(IPaddr string) (uint, error) {
	// Open database
	db, err := maxminddb.Open("data/GeoLite2-City.mmdb")
	if err != nil {
		return 0, err
	}
	defer db.Close()

	// Convert IP astring to net IP format
	ip := net.ParseIP(IPaddr)

	// IP lookup
	var record GeoIPCity
	err = db.Lookup(ip, &record)
	if err == nil {
		fmt.Println(record)
		temp := fmt.Sprintf("%s, %s, %s, %d", record.Country.IsoCode, record.Continent.Code, record.Subdivisions[0].IsoCode, record.Subdivisions[0].GeoNameID)
		fmt.Println(temp)
		return record.Subdivisions[0].GeoNameID, nil
	}

	return 0, err
}

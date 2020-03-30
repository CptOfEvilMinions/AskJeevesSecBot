package geoip

import (
	"fmt"
	"net"

	"github.com/CptOfEvilMinions/AskJeevesSecBot/pkg/config"
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

// InitGeoIPReader input:
// InitGeoIPReader output:
func InitGeoIPReader(cfg *config.Config) (*maxminddb.Reader, error) {
	// Open database
	db, err := maxminddb.Open(cfg.GeoIP.FilePath)
	if err != nil {
		return nil, err
	}
	//defer db.Close()

	return db, nil
}

// IPaddrLocationLookup input: Takes in an IP address as a string
// IPaddrLocationLookup output: Returns Geo lcoation UID
func IPaddrLocationLookup(db *maxminddb.Reader, IPaddr string) (string, uint, error) {
	// Convert IP astring to net IP format
	ip := net.ParseIP(IPaddr)

	// IP lookup
	var record GeoIPCity
	err := db.Lookup(ip, &record)
	if err == nil {
		fmt.Println(record)
		tempLoc := fmt.Sprintf("%s, %s, %s", record.Subdivisions[0].Names["en"], record.Country.IsoCode, record.Continent.Code)
		fmt.Println(tempLoc)
		return tempLoc, record.Subdivisions[0].GeoNameID, nil
	}

	return "", 0, err
}

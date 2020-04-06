package geoip

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/CptOfEvilMinions/AskJeevesSecBot/pkg/config"
	"github.com/oschwald/maxminddb-golang"
)

// https://medium.com/@skdomino/taring-untaring-files-in-go-6b07cf56bc07
func unzipGeoIPdatabase(src string, dst string) error {

	// Opening tar.gz
	var r io.Reader
	r, err := os.Open(src)

	// Create GZIP reader
	gzr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	defer gzr.Close()
	fmt.Println("hello1")

	// Create TAR reader
	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()

		switch {

		// if no more files are found return
		case err == io.EOF:
			return nil

		// return any other error
		case err != nil:
			return err

		// if the header is nil, just skip it (not sure how this happens)
		case header == nil:
			continue

		// If the file is not the database SKIP it
		case header.Name[len(header.Name)-5:] != ".mmdb":
			continue
		}

		// the target location where the dir/file should be created
		fileName := strings.Split(header.Name, "/")
		target := filepath.Join(dst, fileName[1])

		// the following switch could also be done using fi.Mode(), not sure if there
		// a benefit of using one vs. the other.
		// fi := header.FileInfo()

		// check the file type
		switch header.Typeflag {

		// if its a dir and it doesn't exist create it
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}

		// if it's a file create it
		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			// copy over contents
			if _, err := io.Copy(f, tr); err != nil {
				return err
			}

			// manually close here after each file operation; defering would cause each file close
			// to wait until all operations have completed.
			f.Close()
		}
	}

	// Delete tar.gz
	err = os.Remove(src)
	if err != nil {
		return err
	}

	return nil

}

// https://golangcode.com/download-a-file-from-a-url/
// https://dev.maxmind.com/geoip/geoipupdate/
func donwloadGeoIPdatabase(cfg *config.Config) error {
	// Format GeoIP download string
	url := fmt.Sprintf(cfg.GeoIP.URL, cfg.GeoIP.LicenseKey)

	fmt.Printf("[+] - %s - Downloading GeoIP database\n", time.Now().Format(time.RubyDate))

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	//out, err := os.Create(cfg.GeoIP.FilePath)
	out, err := os.Create("data/GeoLite2-City.tar.gz")
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the response to disk
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	// Unzip file
	err = unzipGeoIPdatabase("data/GeoLite2-City.tar.gz", "data/")
	if err != nil {
		return err
	}

	return err
}

// InitGeoIPReader input: config
// InitGeoIPReader output: Return GeoIP database connector
func InitGeoIPReader(cfg *config.Config) (*maxminddb.Reader, error) {
	// Check if file exists
	_, err := os.Stat(cfg.GeoIP.FilePath)
	if os.IsNotExist(err) {
		err := donwloadGeoIPdatabase(cfg)
		if err != nil {
			return nil, err
		}
	}

	// Open database
	db, err := maxminddb.Open(cfg.GeoIP.FilePath)
	if err != nil {
		return nil, err
	}
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

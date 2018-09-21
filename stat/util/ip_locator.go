package util

import (
	"fmt"
	"log"
	"net"
	"path"

	"github.com/oschwald/geoip2-golang"

	"github.com/KyberNetwork/reserve-data/common"
)

const (
	geoDBFile = "GeoLite2-Country.mmdb"

	// UnknownCountry is returned when IP locator can't find the country of IP in database
	UnknownCountry = "unknown"
)

// IPLocator is a resolver that query data of IP from MaxMind's GeoLite2 database.
type IPLocator struct {
	r *geoip2.Reader
}

// NewIPLocator returns an instance of ipLocator.
func NewIPLocator() (*IPLocator, error) {
	dbPath := path.Join(common.CurrentDir(), geoDBFile)
	r, err := geoip2.Open(dbPath)
	if err != nil {
		return nil, err
	}
	return &IPLocator{r: r}, nil
}

// IPToCountry returns the country of given IP address.
// When country of given IP does not exists in database, it returns unknown.
func (il *IPLocator) IPToCountry(ip string) (string, error) {
	IPParsed := net.ParseIP(ip)
	if IPParsed == nil {
		return "", fmt.Errorf("invalid ip %s", ip)
	}
	record, err := il.r.Country(IPParsed)
	if err != nil {
		log.Printf("failed to query data from geo-database!")
		return "", err
	}

	country := record.Country.IsoCode //iso code of country
	if country == "" {
		log.Printf("Can't find country of the given ip: %s", ip)
		country = UnknownCountry
	}
	return country, nil
}

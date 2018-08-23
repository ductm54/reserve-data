package util

import (
	"fmt"
	"log"
	"net"
	"path"

	"github.com/oschwald/geoip2-golang"

	"github.com/KyberNetwork/reserve-data/common"
)

const geoDBFile = "GeoLite2-Country.mmdb"

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
		return "", fmt.Errorf("Can't find country of the given ip: %s", ip)
	}
	return country, nil
}

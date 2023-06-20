package aws

import (
	"errors"
)

var (
	errIPNotFound = errors.New("ip not found")
)

// GetIP returns instance IP to connect
func GetIP(id string, servers []Server, iptype string) (string, error) {
	for _, s := range servers {
		if s.Values["instance-id"] == id {
			switch iptype {
			case "private":
				return s.Values["private-ip"], nil
			default:
				return s.Values["public-ip"], nil
			}
		}
	}
	return "", errIPNotFound
}

package aws

import (
	"errors"
)

var (
	errIPNotFound = errors.New("ip not found")
)

// GetIP returns instance IP to connect
func GetIP(id string, instances []Instance, iptype string) (string, error) {
	for _, instance := range instances {
		if instance.Values["instance-id"] == id {
			switch iptype {
			case "private":
				return instance.Values["private-ip"], nil
			default:
				return instance.Values["public-ip"], nil
			}
		}
	}
	return "", errIPNotFound
}

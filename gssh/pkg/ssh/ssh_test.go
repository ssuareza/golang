package ssh

import (
	"testing"
)

func TestGetPublicKey(t *testing.T) {
	file := "../../testdata/id_rsa"
	_, err := GetPublicKey(file)
	if err != nil {
		t.Error("not able to read public key file")
	}
}

package format

import (
	"os"
	"testing"
)

func TestYAMLToJSON(t *testing.T) {
	// open file
	f, err := os.ReadFile("../testdata/test.yaml")
	if err != nil {
		t.Error("file open failed")
	}

	// convert to json
	j, err := YAMLToJSON(f)
	if err != nil || j == nil {
		t.Error("yaml convertion failed", err)
	}
}

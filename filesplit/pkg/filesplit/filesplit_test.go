package filesplit

import (
	"fmt"
	"os"
	"testing"
)

func TestSplit(t *testing.T) {
	// create test file
	tmp := "/tmp/file.dat"
	size := int64(2097152) // 2MB
	if err := createTestFile(tmp, size); err != nil {
		t.Error(err)
	}

	// and chunk it!
	file, err := Split(tmp)
	if err != nil {
		t.Errorf("not able to split file: %s\n", err)
	}

	expected := 4
	if len(file.Chunks) != expected {
		t.Errorf("chunks number is %v and should be %v", len(file.Chunks), expected)
	}
}

// createTestFile creates a file for testing.
func createTestFile(file string, size int64) error {
	fd, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("failed to create %s", file)
	}
	_, err = fd.Seek(int64(size)-1, 0)
	if err != nil {
		return fmt.Errorf("failed to seek")
	}
	_, err = fd.Write([]byte{0})
	if err != nil {
		return fmt.Errorf("write failed")
	}
	err = fd.Close()
	if err != nil {
		return fmt.Errorf("failed to close file")
	}

	return nil
}

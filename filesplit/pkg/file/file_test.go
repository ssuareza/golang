package file

import "testing"

func TestNew(t *testing.T) {
	// and chunk it!
	file, err := New("../testdata/test.txt")
	if err != nil {
		t.Errorf("not able to split file: %s\n", err)
	}

	expected := 1
	if len(file.Chunks) != expected {
		t.Errorf("chunks number is %v and should be %v", len(file.Chunks), expected)
	}
}

func TestSave(t *testing.T) {
	if err := Save("/tmp/test.txt", []byte{}); err != nil {
		t.Error(err)
	}
}

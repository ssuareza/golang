package rename

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanFilename(t *testing.T) {
	var movies = []struct {
		name     string
		expected string
	}{
		{"Joker.2019.720p.BluRay.x264-[YTS.LT].avi", "Joker 2019 720p BluRay x264 YTS LT"},
		{"Avengers Endgame.2019.720p.BluRay.x264-[YTS.LT].avi", "Avengers Endgame 2019 720p BluRay x264 YTS LT"},
	}

	for _, test := range movies {
		assert.Equal(t, clean(test.name), test.expected)
	}
}

func TestMoveFile(t *testing.T) {
	fileName := "/tmp/TestMovie (2019).avi"
	file, err := os.Create(fileName)
	if err != nil {
		t.Error(err)
	}
	file.Close()

	if err := moveFile(fileName, "/tmp/TestMovie (2019)"); err != nil {
		t.Error(err)
	}

	_, err = os.Stat("/tmp/TestMovie (2019)/TestMovie (2019).avi")
	if os.IsNotExist(err) {
		t.Error("file test does not exists")
	}

	// cleaning files
	if err = os.RemoveAll("./TestMovie (2019)/"); err != nil {
		t.Error("not able to delete test directory")
	}

}

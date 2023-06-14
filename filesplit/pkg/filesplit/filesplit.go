package filesplit

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	fileChunk   = 524288   // 512 KB
	fileMaxSize = 52428800 // 50MB
)

// Chunk contains the file split.
type Chunk struct {
	Key   string
	Value []byte
}

// File represents the splitted file.
type File struct {
	Name   string
	Index  string
	Chunks []Chunk
}

// Split splits a file in smaller pieces.
func Split(file string) (*File, error) {
	// open file
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// get file size
	fileInfo, _ := f.Stat()
	fileSize := fileInfo.Size()

	// files larger than fileMaxSize should be rejected
	if fileSize > fileMaxSize {
		return nil, fmt.Errorf("sorry, we don't process files larger than %v bytes", fileMaxSize)
	}

	// calculate total number of parts the file will be chunked into
	chunksNumber := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

	// split file
	var chunks []Chunk
	var index []string
	for i := uint64(0); i < chunksNumber; i++ {
		chunkSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		chunkBuffer := make([]byte, chunkSize)

		f.Read(chunkBuffer)

		// set name of chunk
		chunkName := fmt.Sprintf("%s:%s", filepath.Base(file), strconv.FormatUint(i, 10))

		if err != nil {
			return nil, err
		}

		// store chunk in struct
		chunks = append(chunks, Chunk{Key: chunkName, Value: chunkBuffer})

		// add to index
		index = append(index, chunkName)
	}

	return &File{Name: file, Index: strings.Join(index, " "), Chunks: chunks}, nil
}

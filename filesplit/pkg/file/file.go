package file

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

var (
	errMaxSize error = fmt.Errorf("sorry, we don't process files larger than %v bytes", fileMaxSize)
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

// New creates a new File.
func New(file string) (*File, error) {
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
		return nil, errMaxSize
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

// Save saves the file with the given content.
func Save(file string, content []byte) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(content)
	if err != nil {
		return err
	}

	return nil
}

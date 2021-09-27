package core

import (
	"compress/flate"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/pedreviljoen/go-flate/common"
)

func calculateFileStats(file string) *common.CompressStats {
	fi, err := os.Stat(file)
	if err != nil {
		return nil
	}

	return &common.CompressStats{
		Size: fi.Size(),
	}
}

func fileExists(filename string) (bool, error) {
	_, err := os.Stat(filename)
	if errors.Is(err, os.ErrNotExist) {
		return false, err
	}

	return true, nil
}

func Compress(filenameOriginal string, filenameNew string, compressionLevel int) (*common.CompressionResult, error) {
	/* Check if it is needed to obtain file path, if so... Can use the below:

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return nil, err
	}
	*/
	fmt.Println("Entering compression with: ", filenameOriginal)

	// Check if file exists
	_, err := fileExists(filenameOriginal)
	if err != nil {
		return nil, err
	}

	// Open file and check before compression stats
	inputFile, err := os.Open(filenameOriginal)
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()
	beforeCompress := calculateFileStats(filenameOriginal)

	// Compress and obtain after compression stats
	outputFile, err := os.Create(filenameNew)
	if err != nil {
		return nil, err
	}
	defer outputFile.Close()

	flateWriter, err := flate.NewWriter(outputFile, flate.BestCompression)
	if err != nil {
		return nil, err
	}
	defer flateWriter.Close()
	io.Copy(flateWriter, inputFile)
	flateWriter.Flush()
	afterCompress := calculateFileStats(filenameNew)

	// Return result
	result := &common.CompressionResult{
		BeforeStats: *beforeCompress,
		AfterStats:  *afterCompress,
	}

	return result, nil
}

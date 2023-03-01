package core

import (
	"compress/flate"
	"io"
	"os"

	"github.com/pedreviljoen/goflate/common"
)

func Compress(filenameOriginal string, filenameNew string, compressionLevel int) (*common.CompressionResult, error) {
	// Check if file exists
	_, err := common.FileExists(filenameOriginal)
	if err != nil {
		return nil, err
	}

	// Open file and check before compression stats
	inputFile, err := os.Open(filenameOriginal)
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()
	beforeCompress := common.CalculateFileStats(filenameOriginal)

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
	afterCompress := common.CalculateFileStats(filenameNew)

	// Return result
	result := &common.CompressionResult{
		BeforeStats: *beforeCompress,
		AfterStats:  *afterCompress,
	}

	return result, nil
}

func Decompress(filenameOriginal string, filenameNew string) (*common.CompressionResult, error) {
	/* Check if it is needed to obtain file path, if so... Can use the below:

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return nil, err
	}
	*/
	// Check if file exists
	_, err := common.FileExists(filenameOriginal)
	if err != nil {
		return nil, err
	}

	// Open file and check before compression stats
	inputFile, err := os.Open(filenameOriginal)
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()
	before := common.CalculateFileStats(filenameOriginal)

	// Compress and obtain after compression stats
	outputFile, err := os.Create(filenameNew)
	if err != nil {
		return nil, err
	}
	defer outputFile.Close()

	flateReader := flate.NewReader(inputFile)
	defer flateReader.Close()
	io.Copy(outputFile, flateReader)

	after := common.CalculateFileStats(filenameNew)

	// Return result
	result := &common.CompressionResult{
		BeforeStats: *before,
		AfterStats:  *after,
	}

	return result, nil
}

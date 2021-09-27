package common

type CompressStats struct {
	Size int64
}

type CompressionResult struct {
	BeforeStats CompressStats
	AfterStats  CompressStats
}

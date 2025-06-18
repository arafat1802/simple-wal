package wal

import "time"

type WalConfig struct {
	WalPath string
	WalMaxNum int
	WalMaxAge time.Duration
	WalSize int64
	WalSegments int
}

func CreateDefaultWalConfig(directory string) WalConfig {
	return WalConfig{
		WalPath: directory,
		WalSize: 16 * 1024 * 1024,
		WalMaxNum: 100,
		WalMaxAge: 24 * time.Hour,
		WalSegments: 1,
	}
}

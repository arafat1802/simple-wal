package wal

import "time"

type WalConfig struct {
	Directory    string
	EnableFsync  bool
	MaxFileSize  uint64
	MaxSegments  int
	SyncInterval time.Duration
}

func CreateDefaultWalConfig(directory string) WalConfig {
	return WalConfig{
		Directory:    directory,
		MaxFileSize:  16 * 1024 * 1024,
		EnableFsync:  true,
		MaxSegments:  1,
		SyncInterval: 300 * time.Millisecond,
	}
}

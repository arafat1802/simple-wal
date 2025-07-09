package wal

import (
	"bufio"
	"context"
	"os"
	"time"
)

// type WalConfig struct {
// 	Directory string
// 	EnableFsync int
// 	MaxFileSize uint64
// 	MaxSegments int
// 	SyncInterval time.Duration
// }

type WriteAheadLog struct {
	directory string
	currFile  *os.File
	bufWriter *bufio.Writer
	syncTimer *time.Timer
	context   context.Context
}

func InitWAL(cfg WalConfig) (*WriteAheadLog, error) {

	os.MkdirAll(cfg.Directory, 0755)

	file, err := os.Create(cfg.Directory + "/wal.log")
	if err != nil {
		panic(err)
	}

	walog := WriteAheadLog{
		directory: cfg.Directory,
		currFile:  file,
		bufWriter: bufio.NewWriter(file),
		syncTimer: time.NewTimer(cfg.SyncInterval),
		context:   context.Background(),
	}

	go walog.syncPeriodically()

	return &walog, nil
}

func serialize() {
	// TODO: implement

}
func WriteEntryWithCheckpoint() {
	// TODO: implement
}
func (walog *WriteAheadLog) WriteEntry(data []byte) error {
	// TODO: implement
	// file, err := os.OpenFile(walog.directory+"/wal.log", os.O_APPEND, 0755)
	// if err != nil {
	// 	panic(err)
	// }

	// file.Write([]byte("hello world"))
	// file.Close()

	walog.bufWriter.Write(data)

	return nil
}

func (walog *WriteAheadLog) FlushAndClose() {
	walog.bufWriter.Flush()
	walog.currFile.Close()
}

// TODO: implement

func ReadAllEntries() {
	// TODO: implement

}

func (walog *WriteAheadLog) syncPeriodically() {
	for {
		select {
		case <-walog.syncTimer.C:
			walog.FlushAndClose()

		case <-walog.context.Done():
			return
		}
	}
}

package wal

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"simple-wal/protobuf"
	"time"

	"google.golang.org/protobuf/proto"
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
func (walog *WriteAheadLog) WriteEntry(entry *protobuf.WalEntry) error {

	data, err := proto.Marshal(entry)
	fmt.Printf("Serialized bytes: %s\n", data)
	walog.bufWriter.Write(data)

	return err
}

func (walog *WriteAheadLog) FlushAndClose() {
	walog.bufWriter.Flush()
	walog.currFile.Close()
}

func (walog *WriteAheadLog) ReadAllEntries(filename string) (*protobuf.WalEntry, error) {
	// Open file in read-only mode
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read all data from the file
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	size := fileInfo.Size()
	data := make([]byte, size)

	_, err = file.Read(data)
	if err != nil {
		return nil, err
	}

	// Unmarshal into WalEntry
	entry := &protobuf.WalEntry{}
	if err := proto.Unmarshal(data, entry); err != nil {
		return nil, err
	}

	return entry, nil
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

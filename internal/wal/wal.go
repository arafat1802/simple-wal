package wal

import (
	"bufio"
	"context"
	"io/ioutil"
	"os"
	"simple-wal/bazel-bin/proto/protobuf_go_proto_/simple-wal/protobuf"
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
	walog.bufWriter.Write(data)

	return err
}

func (walog *WriteAheadLog) FlushAndClose() {
	walog.bufWriter.Flush()
	walog.currFile.Close()
}

// TODO: implement

func (walog *WriteAheadLog) ReadAllEntries(filename string) (*protobuf.WalEntry, error) {
	
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

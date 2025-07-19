package main

import (
	"encoding/json"
	"fmt"
	"simple-wal/internal/wal"
	"simple-wal/protobuf"
	"time"
)

func main() {
	
	cfg := wal.CreateDefaultWalConfig("/home/arafat/")
	walog, err := wal.InitWAL(cfg)

	if err != nil {
		fmt.Printf("failed to init wal: %v\n", err)
		return
	}
	payload := []byte("Hello Arafat")
	
	crc := simpleCRC32(payload)
	isCheckpoint := true
	

	entry1 := &protobuf.WalEntry{
		LogSeqNumber: 123,      // increasing log number
		Data:         json.Marshal(payload),  // raw byte data
		CRC:          crc,      // calculated CRC
		IsCheckPoint: &isCheckpoint,
	}

	walog.WriteEntry(entry1)
	//walog.FlushAndClose()

	time.Sleep(400 * time.Millisecond)
	// entries, err := walog.ReadAllEntries("/home/arafat/wal.log")

	// for  entry := range entries {
	// 	fmt.Println(entry)
	// }

}

func simpleCRC32(data []byte) uint32 {
	var sum uint32 = 0
	for _, b := range data {
		sum += uint32(b)
	}
	return sum
}

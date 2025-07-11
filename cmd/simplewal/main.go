package main

import (
	"fmt"
	"simple-wal/internal/wal"
	"simple-wal/protobuf"
	"time"
)

func main() {
	// wal, _ := wal.InitWAL(wal.CreateDefaultWalConfig("/home/arafat/Desktop/simple-wal"))
	// fmt.Fprint(wal)
	// if err := wal.WriteEntryWithCheckpoint([]byte("Hello World")); err != nil {
	// 	fmt.Printf("failed to write entry: %v\n", err)
	// 	return
	// }
	// if err := wal.WriteEntry([]byte("Hello World")); err != nil {
	// 	fmt.Printf("failed to write entry: %v\n", err)
	// 	return
	// }
	// time.Sleep(301 * time.Millisecond)

	// entries, _ := wal.ReadAllEntries()
	// for _, entry := range entries {
	// 	fmt.Printf("entry: %v\n", string(entry))
	// }
	// wal.Close()

	cfg := wal.CreateDefaultWalConfig("/home/arafat/")
	walog, err := wal.InitWAL(cfg)

	if err != nil {
		fmt.Printf("failed to init wal: %v\n", err)
		return
	}
	entry1 := &protobuf.WalEntry{
		Message: "Hello WAL",
	}
	entry2 := &protobuf.WalEntry{
		Message: "Hello WAL",
	}
	walog.WriteEntry(entry1)
	walog.WriteEntry(entry2)
	//walog.FlushAndClose()

	time.Sleep(400 * time.Millisecond)

	walog.ReadAllEntries("/home/arafat/wal.log")

}

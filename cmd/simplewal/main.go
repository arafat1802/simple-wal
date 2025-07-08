package main

import (
	"fmt"
	"simple-wal/internal/wal"
	"time"
)

func main() {
	wal, _ := wal.InitWAL(wal.CreateDefaultWalConfig("/home/arafat/Desktop/simple-wal"))
	if err := wal.WriteEntryWithCheckpoint([]byte("Hello World")); err != nil {
		fmt.Printf("failed to write entry: %v\n", err)
		return
	}
	if err := wal.WriteEntry([]byte("Hello World")); err != nil {
		fmt.Printf("failed to write entry: %v\n", err)
		return
	}
	time.Sleep(301 * time.Millisecond)

	entries, _ := wal.ReadAllEntries()
	for _, entry := range entries {
		fmt.Printf("entry: %v\n", string(entry))
	}
	wal.Close()
}	

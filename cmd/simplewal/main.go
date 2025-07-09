package main

import (
	"fmt"
	"simple-wal/internal/wal"
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

	walog.WriteEntry([]byte("Hello World\n"))
	walog.WriteEntry([]byte("Hello Arafat\n"))
	//walog.FlushAndClose()

	time.Sleep(400 * time.Millisecond)

}

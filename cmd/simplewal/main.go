package main

import (
	"simple-wal/internal/wal"
)

func main() {
	wal, _ = wal.InitWAL(wal.CreateDefaultWalConfig("/home/arafat/Desktop/simple-wal"))
}

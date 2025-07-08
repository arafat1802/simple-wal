package wal

import (
	"fmt"
	"os"
	"path/filepath"
)

func InitWAL(cfg WALConfig) (*WAL, error) {
	if cfg.Directory == "" {
		return nil, fmt.Errorf("wal: directory must be specified")
	}
	if cfg.FileSize == 0 {
		return nil, fmt.Errorf("wal: file size must be specified")
	}
	if cfg.MaxFiles == 0 {
		return nil, fmt.Errorf("wal: max files must be specified")
	}

	if err := os.MkdirAll(cfg.Directory, 0755); err != nil {
		return nil, fmt.Errorf("wal: failed to create directory: %w", err)
	}

	// don't need this becaule we are working on single file rather than segmentation
	

	// files, err := filepath.Glob(filepath.Join(cfg.Directory, "*.wal"))
	// if err != nil {
	// 	return nil, fmt.Errorf("wal: failed to list files: %w", err)
	// }

	// if len(files) > 0 {
	// 	lastSegmentID, err = FindLastSegmentIndexFiles(files)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("wal: failed to find last segment index: %w", err)
	// 	}
	// } else {
	// 	// create the first log segment
	// 	file, err := CreateSegmentFile(directory, 0)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("wal: failed to create first segment file: %w", err)
	// 	}
	// 	lastSegmentID = 0
	// }
	
	segmentID, err := CreateSegmentFile(cfg.Directory)
	
	if err != nil {
		return nil, fmt.Errorf("wal: failed to create segment file: %w", err)
	}

	

	

}

func serialize() {
	// TODO: implement

}
func WriteEntryWithCheckpoint() {
	// TODO: implement
}
func WriteEntry() {
	// TODO: implement

}

func ReadAllEntries() {
	// TODO: implement

}

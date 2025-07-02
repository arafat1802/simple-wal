package wal

func CreateSegmentFile(dir string) (int, error)) {
	// TODO
	files, err := filepath.Glob(filepath.Join(dir, "*.wal")
	if err != nil {
		return -1, err
	}

	// if no segment is present, create a new one
	if len(files) == 0 {
		file, err := CreateSegmentFile(dir, 0)
		if err != nil {
			return -1, nill
		} else {

		}
		return 0, nil
	}
}

func CreateSegmentFile(dir string, segmentID int) (*os.file, error) {
	// TODO
	file, err := os.Create(filepath.Join(dir, fmt.Sprintf("%d.wal", index)))
	if err != nil {
		return -1, err
	}
	return index, nil
}
if len(files) == 0 {
		file, err := createSegmentFile(dir, 0)
		if err != nil {
			return -1, err
		}
		if err := file.Close(); err != nil {
			return -1, err
		}
		return 0, nil
	}

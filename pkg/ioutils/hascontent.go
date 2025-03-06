package ioutils

import (
	"fmt"
	"io"
)

func getLength(rs io.ReadSeeker) (int64, error) {
	current, err := rs.Seek(0, io.SeekCurrent)
	if err != nil {
		return 0, err
	}
	length, err := rs.Seek(0, io.SeekEnd)
	if err != nil {
		return 0, err
	}
	_, err = rs.Seek(current, io.SeekStart)
	if err != nil {
		return 0, err
	}
	return length, nil
}

func HasContent(rs io.ReadSeeker) (bool, error) {
	if rs == nil {
		return false, fmt.Errorf("reader is nil")
	}
	length, err := getLength(rs)
	if err != nil {
		return false, err
	}
	return length > 2, nil
}

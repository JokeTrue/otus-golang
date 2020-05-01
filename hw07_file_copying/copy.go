package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath string, toPath string, offset, limit int64) error {
	info, err := os.Stat(fromPath)
	if err != nil {
		return err
	}
	if info.IsDir() || info.Size() == 0 {
		return ErrUnsupportedFile
	}
	if offset > info.Size() {
		return ErrOffsetExceedsFileSize
	}

	oldFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer oldFile.Close()

	seekPosition, err := oldFile.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	barLimit := info.Size() - seekPosition
	if limit < barLimit {
		barLimit = limit
	}
	bar := pb.Simple.Start64(barLimit)
	defer bar.Finish()

	bufferLimit := limit
	if limit == 0 {
		bufferLimit = info.Size()
	}
	reader := io.LimitReader(oldFile, bufferLimit)
	barReader := bar.NewProxyReader(reader)

	newFile, err := os.Create(toPath)
	if err != nil {
		return err
	}

	if _, err = io.Copy(newFile, barReader); err != nil {
		return err
	}

	return nil
}

package main

import (
	"errors"
	"io"
	"os"
	"time"

	"github.com/cheggaaa/pb"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrLimits                = errors.New("incorrect offset or limit value")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	var l int64

	if offset < 0 || limit < 0 {
		return ErrLimits
	}

	inFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer inFile.Close()

	inF, err := inFile.Stat()
	if err != nil {
		return ErrUnsupportedFile
	}

	if offset > inF.Size() {
		return ErrOffsetExceedsFileSize
	}

	outFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = inFile.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	if limit == 0 || limit > inF.Size() {
		l = inF.Size()
	} else {
		l = limit
	}

	bar := pb.New(int(l)).SetUnits(pb.U_BYTES).SetRefreshRate(time.Millisecond * 1)
	bar.ShowSpeed = true
	bar.Start()
	reader := bar.NewProxyReader(inFile)

	if _, err := io.CopyN(outFile, reader, l); err != nil && !errors.Is(err, io.EOF) {
		return err
	}

	bar.Finish()
	return nil
}

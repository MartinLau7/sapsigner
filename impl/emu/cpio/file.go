package cpio

import (
	"bytes"
	"io"
	"io/fs"
)

type File struct {
	fi   *FileInfo
	data *bytes.Reader
}

var _ fs.File = (*File)(nil)

func (f *File) Stat() (fs.FileInfo, error) {
	return f.fi, nil
}

func (f *File) Read(bytes []byte) (int, error) {
	return f.data.Read(bytes)
}

func (f *File) Close() error {
	return nil
}

var _ io.ReaderAt = (*File)(nil)

func (f *File) ReadAt(p []byte, off int64) (n int, err error) {
	return f.data.ReadAt(p, off)
}

var _ io.Seeker = (*File)(nil)

func (f *File) Seek(offset int64, whence int) (int64, error) {
	return f.data.Seek(offset, whence)
}

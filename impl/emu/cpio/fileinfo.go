package cpio

import (
	"io/fs"
	"time"
)

type FileInfo struct {
	dev      uint32
	ino      uint32
	mode     fs.FileMode
	uid      uint32
	gid      uint32
	nLink    uint32
	rDev     uint32
	mTime    time.Time
	nameSize uint32
	fileSize int64
	name     string
}

var _ fs.FileInfo = (*FileInfo)(nil)

func (fi *FileInfo) Name() string {
	return fi.name
}

func (fi *FileInfo) Size() int64 {
	return fi.fileSize
}

func (fi *FileInfo) Mode() fs.FileMode {
	return fi.mode
}

func (fi *FileInfo) ModTime() time.Time {
	return fi.mTime
}

func (fi *FileInfo) IsDir() bool {
	return fi.mode.IsDir()
}

func (fi *FileInfo) Sys() any {
	return nil
}

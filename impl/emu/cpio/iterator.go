package cpio

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"iter"
	"strconv"
	"time"
)

type iterator struct {
	r io.Reader
}

func NewIterator(r io.Reader) iter.Seq2[string, *File] {
	i := &iterator{
		r: r,
	}
	return i.seq2
}

func (i *iterator) seq2(yield func(string, *File) bool) {
	for {
		f, err := i.next()
		if errors.Is(err, io.EOF) {
			return
		}
		if err != nil {
			panic(err)
		}

		if !yield(f.fi.name, f) {
			return
		}
	}
}

func (i *iterator) next() (*File, error) {
	hBuf := make([]byte, HeaderSize)
	if _, err := io.ReadFull(i.r, hBuf); err != nil {
		return nil, fmt.Errorf("could not read header: %w", err)
	}

	var h Header
	if _, err := binary.Decode(hBuf, binary.BigEndian, &h); err != nil {
		return nil, fmt.Errorf("could not decode header: %w", err)
	}

	if h.Magic != HeaderMagic {
		return nil, fmt.Errorf("magic is not `070707`: %s", h.Magic)
	}

	dev, err := strconv.ParseUint(string(h.Dev[:]), 8, 18)
	if err != nil {
		return nil, fmt.Errorf("could not parse dev: %w", err)
	}

	ino, err := strconv.ParseUint(string(h.Ino[:]), 8, 18)
	if err != nil {
		return nil, fmt.Errorf("could not parse ino: %w", err)
	}

	mode, err := strconv.ParseUint(string(h.Mode[:]), 8, 18)
	if err != nil {
		return nil, fmt.Errorf("could not parse mode: %w", err)
	}

	uid, err := strconv.ParseUint(string(h.UID[:]), 8, 18)
	if err != nil {
		return nil, fmt.Errorf("could not parse uid: %w", err)
	}

	gid, err := strconv.ParseUint(string(h.GID[:]), 8, 18)
	if err != nil {
		return nil, fmt.Errorf("could not parse gid: %w", err)
	}

	nLink, err := strconv.ParseUint(string(h.NLink[:]), 8, 18)
	if err != nil {
		return nil, fmt.Errorf("could not parse nlink: %w", err)
	}

	rDev, err := strconv.ParseUint(string(h.RDev[:]), 8, 18)
	if err != nil {
		return nil, fmt.Errorf("could not parse rdev: %w", err)
	}

	mTime, err := strconv.ParseUint(string(h.MTime[:]), 8, 33)
	if err != nil {
		return nil, fmt.Errorf("could not parse mtime: %w", err)
	}

	nameSize, err := strconv.ParseUint(string(h.NameSize[:]), 8, 18)
	if err != nil {
		return nil, fmt.Errorf("could not parse namesize: %w", err)
	}

	fileSize, err := strconv.ParseUint(string(h.FileSize[:]), 8, 33)
	if err != nil {
		return nil, fmt.Errorf("could not parse filesize: %w", err)
	}

	nBuf := make([]byte, nameSize)
	if _, err := io.ReadFull(i.r, nBuf); err != nil {
		return nil, fmt.Errorf("could not read header: %w", err)
	}

	if !IsValid(h, nBuf) {
		return nil, fmt.Errorf("invalid entry: %s %s", h, nBuf)
	}

	if IsTrailer(h, nBuf) {
		return nil, io.EOF
	}

	data := make([]byte, fileSize)
	if _, err := io.ReadFull(i.r, data); err != nil {
		return nil, fmt.Errorf("could not read data: %w", err)
	}

	fi := &FileInfo{
		dev:      uint32(dev),
		ino:      uint32(ino),
		mode:     fs.FileMode(mode),
		uid:      uint32(uid),
		gid:      uint32(gid),
		nLink:    uint32(nLink),
		rDev:     uint32(rDev),
		mTime:    time.Unix(int64(mTime), 0),
		nameSize: uint32(nameSize),
		fileSize: int64(fileSize),
		name:     string(nBuf[:nameSize-1]),
	}
	f := &File{
		fi:   fi,
		data: bytes.NewReader(data),
	}
	return f, nil
}

package library

import (
	"bytes"
	"compress/bzip2"
	"context"
	"errors"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"path/filepath"
	"slices"

	"github.com/blacktop/go-macho/pkg/xar"
	"howett.net/ranger"

	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/cpio"
)

const (
	pkgURL    = "http://swcdn.apple.com/content/downloads/27/34/041-98128-A_SYPWICN3KH/5dqkl4rqgbsr18yzy61yeie9g3cmjc5hiv/OSXUpd10.9.pkg"
	xarPath   = "Payload"
	bzip2Addr = 0x352F40D5                                                                                                                      // offset of the first `BZh9` header after offset `bz2SR.*.(*io.SectionReader).off - bz2SR.*.(*io.SectionReader).base` when `bz2SR` is seeked to offset 0x99225F27
	cpioAddr  = 0x000003A4                                                                                                                      // offset of the first `070707` magic in `odcBZ2R`
	ckfwPath  = "./System/Library/PrivateFrameworks/CommerceKit.framework/Versions/A/CommerceKit"                                               // 0x99225F20 + 7
	stagPath  = "./System/Library/PrivateFrameworks/CommerceKit.framework/Versions/A/Resources/storeagent"                                      // 0x9991B4C0 + 13
	ccfwPath  = "./System/Library/PrivateFrameworks/CommerceKit.framework/Versions/A/Frameworks/CommerceCore.framework/Versions/A/CommerceCore" // 0x99D46090 + 8
	fpfwPath  = "./System/Library/PrivateFrameworks/CoreFP.framework/Versions/A/CoreFP"                                                         // 0x9ADB31D0 + 14
	icsxPath  = "./System/Library/PrivateFrameworks/CoreFP.framework/Versions/A/CoreFP.icxs"                                                    // 0x9C95EDF0 + 0
)

var (
	artifactPaths = []string{ckfwPath, stagPath, ccfwPath, fpfwPath, icsxPath}
	artifactCount = len(artifactPaths)
)

func Fetch(ctx context.Context) (map[string][]byte, error) {
	pkgURL, err := url.Parse(pkgURL)
	if err != nil {
		return nil, err
	}

	pkgR, err := ranger.NewReader(&ranger.HTTPRanger{
		Client: http.DefaultClient,
		URL:    pkgURL,
	})
	if err != nil {
		return nil, err
	}

	pkgLen, err := pkgR.Length()
	if err != nil {
		return nil, err
	}

	pkgXR, err := xar.NewReader(pkgR, pkgLen)
	if err != nil {
		return nil, err
	}

	var payloadXF *xar.File
	for _, file := range pkgXR.File {
		if file.Name != xarPath {
			continue
		}

		payloadXF = file
		break
	}
	if payloadXF == nil {
		return nil, &fs.PathError{
			Op:   "open",
			Path: xarPath,
			Err:  fs.ErrNotExist,
		}
	}

	bz2SR := payloadXF.OpenRaw()
	if _, err := bz2SR.Seek(bzip2Addr, io.SeekStart); err != nil {
		return nil, err
	}

	bz2MR := io.MultiReader(bytes.NewBufferString("BZh9"), bz2SR)

	odcBZ2R := bzip2.NewReader(bz2MR)
	if _, err := io.ReadFull(odcBZ2R, make([]byte, cpioAddr)); err != nil {
		return nil, err
	}

	artifacts := make(map[string][]byte, artifactCount)
	for name, file := range cpio.NewIterator(odcBZ2R) {
		if slices.Contains(artifactPaths, name) {
			data, err := io.ReadAll(file)
			if err != nil {
				return nil, err
			}

			artifacts[filepath.Base(name)] = data
		}

		if len(artifacts) == artifactCount {
			return artifacts, nil
		}
	}

	errs := make([]error, artifactCount-len(artifacts))
	for _, cpioPath := range artifactPaths {
		if _, ok := artifacts[filepath.Base(cpioPath)]; ok {
			continue
		}

		err := &fs.PathError{
			Op:   "open",
			Path: cpioPath,
			Err:  fs.ErrNotExist,
		}
		errs = append(errs, err)
	}
	return nil, errors.Join(errs...)
}

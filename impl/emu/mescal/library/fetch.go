package library

import (
	"bytes"
	"compress/bzip2"
	"context"
	"io"
	"io/fs"
	"net/http"
	"net/url"

	"github.com/blacktop/go-apfs/pkg/disk/dmg"
	"github.com/blacktop/go-macho/pkg/xar"
	"howett.net/ranger"

	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/cpio"
)

const (
	dmgURL    = "https://updates.cdn-apple.com/2019/cert/041-88140-20191011-250758ef-d633-428d-afa8-7334f518e125/OSXUpd10.9.1.dmg"
	blkName   = "disk image (Apple_HFS : 3)"
	hfsxPath  = "OSXUpd10.9.1.pkg"
	pkgAddr   = 0x0B7EF000
	pkgSize   = 0x17991000 - 0x0B7EF000
	xarPath   = "OSXUpd10.9.1.pkg/Payload"
	bzip2Addr = 0x07AAB308
	cpioAddr  = 0x000151E5
	cpioPath  = "./System/Library/PrivateFrameworks/CommerceKit.framework/Versions/A/CommerceKit"
)

func Fetch(ctx context.Context) ([]byte, error) {
	diskImageURL, err := url.Parse(dmgURL)
	if err != nil {
		return nil, err
	}

	diskImageR, err := ranger.NewReader(&ranger.HTTPRanger{
		Client: http.DefaultClient,
		URL:    diskImageURL,
	})
	if err != nil {
		return nil, err
	}

	diskImageLen, err := diskImageR.Length()
	if err != nil {
		return nil, err
	}

	diskImageSR := io.NewSectionReader(diskImageR, 0, diskImageLen)

	diskImage, err := dmg.NewDMG(diskImageSR)
	if err != nil {
		return nil, err
	}
	defer diskImage.Close()

	block, err := diskImage.GetBlock(blkName)
	if err != nil {
		return nil, err
	}

	pkgSR := io.NewSectionReader(block, int64(pkgAddr), int64(pkgSize))

	pkgXR, err := xar.NewReader(pkgSR, int64(pkgSize))
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
	if _, err := odcBZ2R.Read(make([]byte, cpioAddr)); err != nil {
		return nil, err
	}

	var frameworkRC io.ReadCloser
	for name, file := range cpio.NewIterator(odcBZ2R) {
		if name != cpioPath {
			continue
		}

		frameworkRC = file
		break
	}
	if frameworkRC == nil {
		return nil, &fs.PathError{
			Op:   "open",
			Path: cpioPath,
			Err:  fs.ErrNotExist,
		}
	}
	defer frameworkRC.Close()

	return io.ReadAll(frameworkRC)
}

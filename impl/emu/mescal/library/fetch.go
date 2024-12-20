package library

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/fs"
	"net/http"

	"github.com/blacktop/go-macho/pkg/xar"
	"github.com/korylprince/go-cpio-odc"
)

const (
	pkgURL   = "https://swcdn.apple.com/content/downloads/11/62/041-88157-A_9MV302JSGJ/nut08bz4n2byymwkltjsxcsans8pwndovm/AppStoreUpdate.pkg"
	xarPath  = "Payload"
	cpioPath = "System/Library/PrivateFrameworks/CommerceKit.framework/Versions/A/CommerceKit"
)

func Fetch(ctx context.Context) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, pkgURL, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", res.Status)
	}

	bodyRA, err := newReaderAt(res.Body)
	if err != nil {
		return nil, err
	}

	bodyXR, err := xar.NewReader(bodyRA, res.ContentLength)
	if err != nil {
		return nil, err
	}

	var payloadXF *xar.File
	for _, file := range bodyXR.File {
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
	payloadXF.EncodingMimetype = "application/x-bzip2"

	payloadRC, err := payloadXF.Open()
	if err != nil {
		return nil, err
	}
	defer payloadRC.Close()

	payloadFS, err := cpio.NewFS(payloadRC)
	if err != nil {
		return nil, err
	}

	frameworkRC, err := payloadFS.Open(cpioPath)
	if err != nil {
		return nil, err
	}
	defer frameworkRC.Close()

	return io.ReadAll(frameworkRC)
}

func newReaderAt(r io.Reader) (io.ReaderAt, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(data), nil
}

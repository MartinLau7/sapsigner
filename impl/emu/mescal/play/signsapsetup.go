package play

import (
	"bytes"
	"context"
	"fmt"
	"howett.net/plist"
	"io"
	"net/http"
)

const (
	signSAPSetupURL = "https://play.itunes.apple.com/WebObjects/MZPlay.woa/wa/signSapSetup"
)

// signSAPSetupBody is a Go representation of request and response bodies of signSAPSetupURL.
//
//	<?xml version="1.0" encoding="UTF-8"?>
//	<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
//	<plist version="1.0">
//		<dict>
//			<key>sign-sap-setup-buffer</key>
//			<data>%s</data>
//		</dict>
//	</plist>
type signSAPSetupBody struct {
	SignSAPSetupBuffer []byte `plist:"sign-sap-setup-buffer"`
}

func SignSAPSetup(ctx context.Context, data []byte) ([]byte, error) {
	reqBody := signSAPSetupBody{
		SignSAPSetupBuffer: data,
	}

	var reqBuffer bytes.Buffer
	if err := plist.NewEncoder(&reqBuffer).Encode(reqBody); err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, signSAPSetupURL, &reqBuffer)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-plist")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", res.Status)
	}

	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	resReader := bytes.NewReader(resBytes)

	var resBody signSAPSetupBody
	if err := plist.NewDecoder(resReader).Decode(&resBody); err != nil {
		return nil, err
	}
	setupBuffer := resBody.SignSAPSetupBuffer

	return setupBuffer, nil
}

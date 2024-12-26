package mzinit

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

const (
	signSAPSetupCertURL = "https://s.mzstatic.com/sap/setup.crt" // https://init.itunes.apple.com/WebObjects/MZInit.woa/wa/signSapSetupCert
)

func SignSAPSetupCert(ctx context.Context) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, signSAPSetupCertURL, nil)
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

	return io.ReadAll(res.Body)
}

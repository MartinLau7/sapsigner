package guid

import (
	"fmt"
	"net"
)

func Get() ([]byte, error) {
	nis, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	if nis == nil {
		return nil, fmt.Errorf("no network interface found")
	}

	var nhwa net.HardwareAddr
	for _, ni := range nis {
		switch ni.Name {
		case "en0":
			return ni.HardwareAddr, nil
		case "en1":
			nhwa = ni.HardwareAddr
		}
	}
	if nhwa == nil {
		nhwa = nis[0].HardwareAddr
	}

	return nhwa, nil
}

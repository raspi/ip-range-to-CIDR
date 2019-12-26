package lib

import (
	"fmt"
	"net"
)

func ToCidr(start, end net.IP) (string, uint8, error) {
	if !((start.To4() != nil && end.To4() != nil) || (start.To16() != nil && end.To16() != nil)) {
		return ``, 0, fmt.Errorf(`invalid IP address(es)`)
	}

	if start.To4() != nil {
		start = start.To4()
	} else {
		start = start.To16()
	}

	if end.To4() != nil {
		end = end.To4()
	} else {
		end = end.To16()
	}

	mask := make([]byte, len(start))

	for idx, _ := range start {
		mask[idx] = 255 - (start[idx] ^ end[idx])
	}

	ipnet := net.IPNet{
		IP:   start,
		Mask: mask,
	}

	ones, _ := ipnet.Mask.Size()

	return start.String(), uint8(ones), nil

}

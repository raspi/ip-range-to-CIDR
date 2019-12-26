package lib

import (
	"net"
	"testing"
)

func TestCidr24(t *testing.T) {
	ip, mask, err := ToCidr(net.ParseIP(`192.168.0.0`), net.ParseIP(`192.168.0.255`))
	if err != nil {
		t.Fail()
	}

	if ip != `192.168.0.0` {
		t.Fail()
	}

	if mask != 24 {
		t.Fail()
	}
}

func TestCidr0(t *testing.T) {
	ip, mask, err := ToCidr(net.ParseIP(`0.0.0.0`), net.ParseIP(`255.255.255.255`))
	if err != nil {
		t.Fail()
	}

	if ip != `0.0.0.0` {
		t.Fail()
	}

	if mask != 0 {
		t.Fail()
	}
}

func TestCidr32(t *testing.T) {
	ip, mask, err := ToCidr(net.ParseIP(`1.2.3.4`), net.ParseIP(`1.2.3.4`))
	if err != nil {
		t.Fail()
	}

	if ip != `1.2.3.4` {
		t.Fail()
	}

	if mask != 32 {
		t.Fail()
	}
}

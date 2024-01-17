package otp

import (
	"crypto/sha1"
	"testing"

	//lint:ignore SA1019 only used for testing
	"golang.org/x/crypto/md4" //nolint:staticcheck
)

var BadKeys = []Key{
	{
		Method: "crypto!",
	},
	{
		Method: "totp",
	},
	{
		Method: "totp",
		Label:  "t@w",
	},
	{
		Method:   "totp",
		Label:    "t@w",
		Secret32: "abc123",
	},
	{
		Method:   "totp",
		Label:    "t@w",
		Secret32: "MFRGGZDFMZTWQ2LK",
		Issuer:   "issuer",
		Algo:     md4.New,
	},
	{
		Method:   "totp",
		Label:    "t@w",
		Secret32: "MFRGGZDFMZTWQ2LK",
		Issuer:   "issuer",
		Algo:     sha1.New,
		Digits:   99,
	},
	{
		Method:   "totp",
		Label:    "t@w",
		Secret32: "MFRGGZDFMZTWQ2LK",
		Issuer:   "issuer",
		Algo:     sha1.New,
		Digits:   6,
		Period:   -42,
	},
}

func TestBadKeys(t *testing.T) {
	for _, k := range BadKeys {
		if err := k.Validate(); err == nil {
			t.Errorf("bad Key didn't produce error on Validate(): %v", k)
		}
	}
}

package otp

import (
	"crypto/sha1"
	"hash"
	"testing"
)

func invalidTestHash() hash.Hash {
	return nil
}

func TestBadKeys(t *testing.T) {
	badKeys := []Key{
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
			Algo:     invalidTestHash,
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

	for _, k := range badKeys {
		if err := k.Validate(); err == nil {
			t.Errorf("bad Key didn't produce error on Validate(): %v", k)
		}
	}
}

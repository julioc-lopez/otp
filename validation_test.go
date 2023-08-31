package otp

import (
	"crypto/sha1"
	"testing"

	"golang.org/x/crypto/md4"
)

var BadKeys = []Key{
	Key{
		Method: "crypto!",
	},
	Key{
		Method: "totp",
	},
	Key{
		Method: "totp",
		Label:  "t@w",
	},
	Key{
		Method:   "totp",
		Label:    "t@w",
		Secret32: "abc123",
	},
	Key{
		Method:   "totp",
		Label:    "t@w",
		Secret32: "MFRGGZDFMZTWQ2LK",
		Issuer:   "issuer",
		Algo:     md4.New,
	},
	Key{
		Method:   "totp",
		Label:    "t@w",
		Secret32: "MFRGGZDFMZTWQ2LK",
		Issuer:   "issuer",
		Algo:     sha1.New,
		Digits:   99,
	},
	Key{
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

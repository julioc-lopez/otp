package otp

import (
	"crypto/sha1"
	"testing"
)

func TestNewTOTPKey(t *testing.T) {
	if _, err := NewTOTPKey(
		"label",
		"MFRGGZDFMZTWQ2LK",
		"issuer",
		sha1.New,
		6,
		30,
	); err != nil {
		t.Errorf("failed to build new totp key:\n%v", err)
	}
}

func TestNewBadTotp(t *testing.T) {
	if _, err := NewTOTPKey(
		"label",
		"MifdasfsfdsfFRGGZDFMZTWQ2LK0",
		"issuer",
		sha1.New,
		6,
		30,
	); err == nil {
		t.Fail()
	}
}

func TestNewBadHotp(t *testing.T) {
	if _, err := NewHOTPKey(
		"label",
		"MFRfadfdssdGGZDFMZTWQ2LK",
		"issuer",
		sha1.New,
		6,
		30,
	); err != nil {
		t.Fail()
	}
}

func TestNewHOTPKey(t *testing.T) {
	if _, err := NewHOTPKey(
		"label",
		"MFRGGZDFMZTWQ2LK",
		"issuer",
		sha1.New,
		6,
		30,
	); err != nil {
		t.Error("failed to build new hotp key")
	}
}

func TestNewKeyFromURI(t *testing.T) {
	uri := "otpauth://totp/label?secret=MFRGGZDFMZTWQ2LK&issuer=theIssuer&algo=SHA512"
	if k, err := NewKey(uri); err != nil {
		t.Errorf("Constructor failed:\n%v\n%v", err, k)
	}

	uri = "blahblah"
	if _, err := NewKey(uri); err == nil {
		t.Error("Should have failed...")
	}

	uri = "otpauth://totp/label?secret=MFRGGZDFMZTWQ2LK&issuer=theIssuer&algo=SHA512&period=0"
	if _, err := NewKey(uri); err == nil {
		t.Error("Should have failed...")
	}

}

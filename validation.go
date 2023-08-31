package otp

import (
	"encoding/base32"
	"errors"
	"strings"
)

func (k Key) validateMethod() error {
	if !stringInSlice(k.Method, methods) {
		return errors.New("invalid method value")
	}
	return nil
}

func (k Key) validateLabel() error {
	if len(k.Label) == 0 {
		return errors.New("missing value for label")
	}
	return nil
}

func (k Key) validateSecret32() error {
	if len(k.Secret32) == 0 {
		return errors.New("missing value for secret")
	}

	if _, err := base32.StdEncoding.DecodeString(strings.ToUpper(k.Secret32)); err != nil {
		return errors.New("invalid Base32 value for secret")
	}

	return nil
}

func (k Key) validateAlgo() error {
	if !hashInSlice(k.Algo, Hashes) {
		return errors.New("invalid hashing algorithm")
	}
	return nil
}

func (k Key) validateDigits() error {
	if !(k.Digits == 6 || k.Digits == 8) {
		return errors.New("digit is not equal to 6 or 8")
	}
	return nil
}

func (k Key) validatePeriod() error {
	if k.Method == "totp" && k.Period < 1 {
		return errors.New("period can not have a non-positive value")
	}
	return nil
}

// Validate checks if the key parameters conform to the specification.
// In invalid, an error is returns.
func (k Key) Validate() error {
	for _, v := range []func() error{
		k.validateMethod,
		k.validateLabel,
		k.validateSecret32,
		k.validateAlgo,
		k.validateDigits,
		k.validatePeriod,
	} {
		if err := v(); err != nil {
			return err
		}
	}

	return nil
}

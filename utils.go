package otp

import (
	"encoding/base32"
	"reflect"
	"runtime"
	"strings"
)

// http://stackoverflow.com/a/7053871/3582177
func getFuncName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func hashInSlice(a Hash, list []Hash) bool {
	for _, b := range list {
		if getFuncName(b) == getFuncName(a) {
			return true
		}
	}
	return false
}

var base32NoPadding = base32.StdEncoding.WithPadding(base32.NoPadding)

func decodeBase32(s string) ([]byte, error) {
	s = strings.ToUpper(s)

	if b, err := base32NoPadding.DecodeString(s); err == nil {
		return b, nil
	}

	// re-try allowing padding
	return base32.StdEncoding.DecodeString(s)
}

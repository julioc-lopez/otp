package otp

import (
	"encoding/base32"
	"reflect"
	"runtime"
	"slices"
	"strings"
	"unicode"
)

// http://stackoverflow.com/a/7053871/3582177
func getFuncName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func stringInSlice(a string, list []string) bool {
	return slices.Contains(list, a)
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
	s = strings.Map(func(c rune) rune {
		if c == '-' || unicode.IsSpace(c) {
			return -1
		}

		return unicode.ToUpper(c)
	}, s)

	if b, err := base32.StdEncoding.DecodeString(s); err == nil {
		return b, nil
	}

	// re-try allowing truncated strings with no padding
	return base32NoPadding.DecodeString(s)
}
